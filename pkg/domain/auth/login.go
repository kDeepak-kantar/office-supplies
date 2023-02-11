// Package auth handles user authentication.
// Currently it only supports authentication agails Google SAML service.
package auth

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/Deepak/pkg/storage/db/user"
)

const (
	Adminuser  string = "admin"
	Normaluser string = "normal"
)

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
	userDetails.Role = Adminuser
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

func (r *domain) Scheduler() {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 9, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C

		if time.Now().Weekday() == time.Monday {
			adminEmails, err := r.User.GetAdminEmails()
			if err != nil {
				log.Println(err)
				continue
			}
			for _, adminEmail := range adminEmails {
				if err := r.sendEmail(adminEmail); err != nil {
					log.Println(err)
					continue
				}
			}
		}

	}
}
func (r *domain) sendEmail(to string) error {
	from := "Kurravenkata.saideepak.kantar@blackwoodseven.com"
	subject := "Pending Orders"
	body := "This is a reminder about the pending orders. Please check and take necessary action."

	// Set up authentication information.
	auth := smtp.PlainAuth("", "Kurravenkata.saideepak.kantar@blackwoodseven.com", "Deepak@cse123", "smtp.example.com")
	err := smtp.SendMail(
		"smtp.example.com:25",
		auth,
		from,
		[]string{to},
		[]byte("Subject: "+subject+"\r\n"+body),
	)
	if err != nil {
		return err
	}

	return nil
}
