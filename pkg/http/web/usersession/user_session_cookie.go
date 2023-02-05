package usersession

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Cookie struct {
	UserID         string
	OrganizationID uint
	AuthToken      string
}

func (c *Cookie) encode() ([]byte, error) {
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	payload := base64Encode(out)
	return payload, nil
}

const CookieKey = "NeoUserSession"

type SetCookieRequest struct {
	Context *gin.Context
	UserID  string
}

func (s *service) SetCookie(req *SetCookieRequest) error {
	val := Cookie{
		UserID: req.UserID,
	}

	valByte, err := val.encode()
	if err != nil {
		return err
	}

	// 60 seconds per minute, 60 min in an hour, 24 hours in a day
	maxAge := 60 * 60 * 24 * time.Second
	cookie := &http.Cookie{
		Name:     CookieKey,
		Value:    string(valByte),
		MaxAge:   int(maxAge.Seconds()),
		Path:     "/",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
	}

	req.Context.SetCookie(
		cookie.Name,
		cookie.Value,
		cookie.MaxAge,
		cookie.Path,
		cookie.Domain,
		cookie.Secure,
		cookie.HttpOnly,
	)
	return nil
}

func (s *service) GetCookie(c *gin.Context) (*Cookie, error) {
	cookie, err := c.Cookie(CookieKey)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil, nil
		}
		return nil, err
	}
	if cookie == "" {
		return nil, nil
	}

	decodedCookie, err := base64Decode([]byte(cookie))
	if err != nil {
		return nil, err
	}

	var userCookie Cookie
	err = json.Unmarshal(decodedCookie, &userCookie)
	if err != nil {
		return nil, err
	}
	return &userCookie, nil
}

func (s *service) DeleteCookie(c *gin.Context) {
	c.SetCookie(CookieKey, "", 0, "/", "", true, true)
}

func base64Decode(v []byte) ([]byte, error) {
	decoded := make([]byte, base64.URLEncoding.DecodedLen(len(v)))
	b, err := base64.URLEncoding.Decode(decoded, v)
	if err != nil {
		return nil, err
	}

	return decoded[:b], nil
}

func base64Encode(v []byte) []byte {
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(v)))
	base64.URLEncoding.Encode(encoded, v)
	return encoded
}
