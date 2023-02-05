// Package auth handles user authentication.
// Currently it only supports authentication agails Google SAML service.
package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"

	"github.com/Deepak/pkg/storage/db/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

//const path string = "~/token.json"

type LoginRequest struct {
	Email string
	Token string
}

type LoginRespose struct {
	User *user.User
}

var (
	ErrInvalidUser = errors.New("not a valid user")
)

// LoginUser will validate the user access token from SAML serivce (eg. Google),
// create new user (if it's a first time login - this function also sign's up the user).
func (r *domain) LoginUser(req LoginRequest) (*LoginRespose, error) {
	claims, err := validateGSuiteToken(req.Token)
	if err != nil {
		return nil, err
	}
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/token.json")
	if err != nil || len(b) == 0 {
		err = saveToken()
		fmt.Println("error is", err)
	}

	if claims.Hd != "blackwoodseven.com" {
		return nil, fmt.Errorf("Denied")
	}

	user, err := r.User.Create(claims.Name, claims.Email)
	if user == nil || err != nil {
		return nil, fmt.Errorf("not found")
	}

	return &LoginRespose{
		User: user,
	}, nil
}

func (r *domain) GetAllUsers() ([]*user.User, error) {
	return r.User.GetAll()
}

func (r *domain) GetUserRole(userId string) (string, error) {
	userDetails, err := r.User.GetUserByStringId(userId)
	if userDetails == nil || err != nil {
		return "", ErrInvalidUser
	}
	return userDetails.Role, nil
}

func (r *domain) AdminAccess(userId string) (*user.User, error) {
	userDetails, err := r.User.GetUserByStringId(userId)
	if userDetails == nil || err != nil {
		return nil, ErrInvalidUser
	}
	userDetails.Role = "Admin"
	err = r.User.UpdateUser(userDetails)
	if err != nil {
		return nil, err
	}
	return userDetails, nil
}

func (r *domain) RemoveUser(userId string) error {
	userDetails, err := r.User.GetUserByStringId(userId)
	if userDetails == nil || err != nil {
		return ErrInvalidUser
	}
	err = r.User.RemoveUser(userDetails)
	if err != nil {
		return err
	}
	return nil
}

func saveToken() error {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/credentials.json")
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return fmt.Errorf("unable to read authorization code: %v", err)
	}
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return fmt.Errorf("unable to convert to oauth token: %v", err)
	}
	f, err := os.OpenFile(pwd+"/token.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(tok)
	return nil
}
func (d *domain) Scheduler() {
	admins, err := d.User.GetAdminEmails()
	if err != nil {
		panic(err)
	}
	sender := "kvsdeepak132000@gmail.com"
	password := "Deepakcse123@"

	// Loop through all administrators and send an email
	for _, admin := range admins {
		to := admin

		// Set up the message
		msg := []byte(fmt.Sprintf("To: %s\r\nSubject: Should Approve Orde\r\n\r\nThis is an example email.", to))

		// Connect to the SMTP server and send the email
		err := smtp.SendMail("smtp.example.com:587", smtp.PlainAuth("", sender, password, "smtp.example.com"), sender, []string{to}, msg)
		if err != nil {
			log.Fatalf("Failed to send email to %s: %v", to, err)
		}
		fmt.Printf("Sent email to %s\n", to)
	}
	return
}
