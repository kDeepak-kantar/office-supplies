package user

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(name, email string) (*User, error)
	GetAll() ([]*User, error)
	GetUserByStringId(id string) (*User, error)
	GetUsersByStringIds(ids []string) ([]*User, error)
	UpdateUser(user *User) error
	RemoveUser(user *User) error
	GetAdminEmails() ([]string, error)
}

type service struct {
	db *gorm.DB
}

func Init(connection *gorm.DB) (Repository, error) {
	err := connection.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	return &service{
		db: connection,
	}, nil
}

func (r *service) Create(name, email string) (*User, error) {
	user := User{}

	err := r.db.Preload(clause.Associations).Where(`email = ?`, email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.Name = name
			user.Email = email
			err = r.db.Create(&user).Error
		} else {
			return nil, err
		}
	} else {
		user.Name = name
		err = r.db.Save(&user).Error
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *service) GetAll() ([]*User, error) {
	var userlis []*User
	err := r.db.
		Order("role").
		Order("name").
		Find(&userlis).Error
	if err != nil {
		return nil, err
	}
	return userlis, nil
}

func (r *service) GetUserByStringId(id string) (*User, error) {
	user := User{}
	err := r.db.Preload(clause.Associations).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *service) GetUsersByStringIds(ids []string) ([]*User, error) {
	users := []*User{}
	err := r.db.
		Preload(clause.Associations).
		Where("id in (?)", ids).
		Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}

func (r *service) UpdateUser(user *User) error {
	return r.db.Save(&user).Error
}

func (r *service) RemoveUser(user *User) error {
	return r.db.Debug().Where("id = ?", user.ID).Delete(&user).Error
}

func (r *service) GetAdminEmails() ([]string, error) {
	var emails []string
	var users []*User
	err := r.db.Where("role = ?", "admin").Find(&users).Error
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		emails = append(emails, user.Email)
	}
	return emails, nil
}
