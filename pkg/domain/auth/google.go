package auth

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Deepak/pkg/logger"
)

const (
	gCertsBaseUrl = "https://www.googleapis.com/oauth2/v1/certs"
)

var (
	errMalformedJwt     = errors.New("error: JWT is malformed")
	errInvalidSignature = errors.New("error: invalid signature")
	errUntrustedIssuer  = errors.New("error: issuer is untrusted")
	errExpiredToken     = errors.New("error: token has expired")
)

var client = &http.Client{
	Timeout: time.Second * 30,
}

type GAccountClaims struct {
	Iss           string
	Sub           string
	Azp           string
	Iat           int
	Exp           int64
	Hd            string
	Email         string
	EmailVerified bool `json:"email_verified"`
	Name          string
	Picture       string
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
}

type gJwtHeader struct {
	Alg string
	Kid string
	Typ string
}

type gPublicKey struct {
	id  string
	key *rsa.PublicKey
}

var gPublicKeys []gPublicKey
var gNextFetch time.Time
var gNextFetchMutex sync.Mutex

// Verify the integrity of the ID token
// After you receive the ID token by HTTPS POST, you must verify the integrity of the token.
// To verify that the token is valid, ensure that the following criteria are satisfied:
//
//   - The ID token is properly signed by Google
//   - The value of aud in the ID token is equal to one of your app's client IDs
//   - The value of iss in the ID token is equal to accounts.google.com or
//     https://accounts.google.com.
//   - The expiry time (exp) of the ID token has not passed.
//   - If you want to restrict access to only members of your G Suite domain,
//     verify that the ID token has an hd claim that matches your G Suite domain name.
func validateGSuiteToken(token string) (*GAccountClaims, error) {
	parts := bytes.SplitN([]byte(token), []byte("."), 3)
	if len(parts) != 3 {
		return nil, errMalformedJwt
	}

	decodedHeader, err := decode(parts[0])
	if err != nil {
		return nil, err
	}

	var header gJwtHeader
	if err := json.Unmarshal(decodedHeader, &header); err != nil {
		return nil, err
	}

	if header.Alg != "RS256" {
		return nil, fmt.Errorf("error: unsupported algorithm: %s", header.Alg)
	}

	var publicKey *rsa.PublicKey
	if publicKey = getPublicKey(header.Kid); publicKey == nil {
		return nil, fmt.Errorf("error: invalid header kid: %s", header.Kid)
	}

	h := sha256.New()
	h.Write(parts[0])
	h.Write([]byte("."))
	h.Write(parts[1])
	hash := h.Sum(nil)

	signature, err := decode(parts[2])
	if err != nil {
		return nil, err
	}

	if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, signature); err != nil {
		return nil, errInvalidSignature
	}

	payload, err := decode(parts[1])
	if err != nil {
		return nil, err
	}

	var claims GAccountClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}

	if claims.Iss != "accounts.google.com" && claims.Iss != "https://accounts.google.com" {
		return &claims, errUntrustedIssuer
	}

	exp := time.Unix(claims.Exp, 0)
	if time.Now().After(exp) {
		return &claims, errExpiredToken
	}

	return &claims, nil
}

func updateCerts() error {
	resp, err := client.Get(gCertsBaseUrl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to get certs: %+v", resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var root interface{}
	if err := json.Unmarshal(body, &root); err != nil {
		return err
	}

	certs, ok := root.(map[string]interface{})
	if !ok {
		return fmt.Errorf("Expected map[string]interface{} from response and got: %v", root)
	}

	var keys []gPublicKey

	for kid, val := range certs {
		rawCert, ok := val.(string)
		if !ok {
			logger.Log("", logger.SeverityWarning, "Raw certificate is not a string type", nil, nil)
			continue
		}

		block, _ := pem.Decode([]byte(rawCert))
		if block == nil {
			logger.Log("", logger.SeverityWarning, "Could not decode PEM", nil, nil)
			continue
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			logger.Log("updateCerts",
				logger.SeverityWarning,
				"Could not parse certificate",
				err,
				nil)
			continue
		}

		if key, ok := cert.PublicKey.(*rsa.PublicKey); ok {
			keys = append(keys, gPublicKey{id: kid, key: key})
		}
	}

	cacheControl := resp.Header["Cache-Control"]
Loop:
	for _, control := range cacheControl {
		parts := strings.Split(control, ", ")
		for _, part := range parts {
			if !strings.HasPrefix(part, "max-age=") {
				continue
			}

			parts := strings.Split(part, "=")
			if len(parts) != 2 {
				continue
			}

			if i, err := strconv.Atoi(parts[1]); err == nil {
				deadline := time.Duration(i) * time.Second
				gNextFetch = time.Now().Add(deadline)
				break Loop
			}
		}
	}

	gPublicKeys = keys

	return nil
}

func getPublicKey(kid string) *rsa.PublicKey {
	gNextFetchMutex.Lock()
	defer gNextFetchMutex.Unlock()

	if time.Now().After(gNextFetch) {
		if err := updateCerts(); err != nil {
			logger.Log("getPublicKey",
				logger.SeverityError,
				"Updating certificates failed",
				err,
				nil)
			return nil
		}
	}

	for _, key := range gPublicKeys {
		if key.id == kid {
			return key.key
		}
	}

	return nil
}

func decode(v []byte) ([]byte, error) {
	if rem := len(v) % 4; rem > 0 {
		for i := 0; i < 4-rem; i += 1 {
			v = append(v, byte('='))
		}
	}

	decoded := make([]byte, base64.URLEncoding.DecodedLen(len(v)))
	b, err := base64.URLEncoding.Decode(decoded, v)
	if err != nil {
		return nil, err
	}

	return decoded[:b], nil
}
