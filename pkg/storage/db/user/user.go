package user

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// An User is a person that can later participate on a CoffeeDate.
type User struct {
	ID    *uuid.UUID `gorm:"type:char(36);primary_key"`
	Name  string     `gorm:"column:name;type:varchar(255);not null;"`
	Email string     `gorm:"column:email;type:varchar(255);not null;unique;"`
	Role  string     `gorm:"column:role;type:varchar(255);not null;default:normal;"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == nil {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}

		user.ID = &uuid
	}

	return nil
}
