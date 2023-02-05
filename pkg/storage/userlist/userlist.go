package userlist

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ItemID   string
	Quantity string
	OrderID  string
}

type Order struct {
	gorm.Model
	UserID        *uuid.UUID
	Items         []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RequestedDate string
	DueDate       string
	Status        string
}
