package userlist

import (
	"errors"
	"fmt"
	"time"

	"github.com/Deepak/pkg/storage/db/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateUserList(c *Order) error
	GetAllist() ([]*Order, error)
	GetOrderByStringId(id string) (*Order, error)
	// GetUsersListsByStringIds(ids []string) ([]*Order, error)
	UpdateList(c *Order) error
	GetAllApproved() ([]*Order, error)
	GetAllNotApproved() ([]*Order, error)
	GetOrderByStringiId(id int) (*Order, error)
	SendRemainder() (map[string]interface{}, error)
	GetUserList(id string) ([]*Order, error)
}

type service struct {
	db *gorm.DB
}

func Init(connection *gorm.DB) (Repository, error) {
	err := connection.AutoMigrate(&Item{})
	if err != nil {
		return nil, err
	}
	err = connection.AutoMigrate(&Order{})
	if err != nil {
		return nil, err
	}
	return &service{
		db: connection,
	}, nil
}

func (r *service) CreateUserList(k *Order) error {
	// fmt.Println(k)
	if k.Status == "" {
		k.Status = "pending"
	}
	if err := r.db.Create(&k).Error; err != nil {
		panic(err)
	}

	return nil
}

//	func (r *service) GetAllist() ([]*Order, error) {
//		var usrlists []*Order
//		err := r.db.Preload(clause.Associations).
//			Find(&usrlists).Error
//		if err != nil {
//			return nil, err
//		}
//		return usrlists, nil
//	}
func (r *service) GetAllist() ([]*Order, error) {
	var usrlists []*Order
	err := r.db.Preload(clause.Associations).
		Find(&usrlists).Error
	if err != nil {
		return nil, err
	}

	for _, order := range usrlists {
		order.DueDate = formatDate(order.DueDate)
		order.RequestedDate = formatDate(order.RequestedDate)
	}

	return usrlists, nil
}

func (r *service) GetUserList(id string) ([]*Order, error) {
	var usrlists []*Order
	err := r.db.Preload(clause.Associations).Where("user_id = ?", id).
		Find(&usrlists).Error
	if err != nil {
		return nil, err
	}

	for _, order := range usrlists {
		order.DueDate = formatDate(order.DueDate)
		order.RequestedDate = formatDate(order.RequestedDate)
	}

	return usrlists, nil
}

func formatDate(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("02/01/2006")
}

func (r *service) GetAllNotApproved() ([]*Order, error) {
	var usrlists []*Order
	err := r.db.Preload(clause.Associations).
		Where("status = ?", "pending").
		Find(&usrlists).Error
	if err != nil {
		return nil, err
	}
	for _, order := range usrlists {
		order.DueDate = formatDate(order.DueDate)
		order.RequestedDate = formatDate(order.RequestedDate)
	}
	return usrlists, nil
}
func (r *service) GetAllApproved() ([]*Order, error) {
	var usrlists []*Order
	err := r.db.Preload(clause.Associations).
		Where("status = ?", "approved").
		Find(&usrlists).Error
	if err != nil {
		return nil, err
	}
	for _, order := range usrlists {
		order.DueDate = formatDate(order.DueDate)
		order.RequestedDate = formatDate(order.RequestedDate)
	}
	return usrlists, nil

}

func (r *service) GetOrderByStringId(id string) (*Order, error) {
	userlist := Order{}
	err := r.db.Preload(clause.Associations).Where("id = ?", id).First(&userlist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userlist, nil
}
func (r *service) UpdateList(c *Order) error {
	return r.db.Save(&c).Error
}

func (r *service) GetOrderByStringiId(id int) (*Order, error) {
	userlist := Order{}
	err := r.db.Preload(clause.Associations).Where("id = ?", id).First(&userlist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userlist, nil
}

func (r *service) SendRemainder() (map[string]interface{}, error) {
	var orders []Order
	err := r.db.Where("status = ?", "pending").Preload(clause.Associations).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	today := time.Now().UTC()
	reminders := []string{}
	for _, order := range orders {
		dueDate, err := time.Parse("02/01/2006", order.DueDate)
		if err != nil {
			return nil, err
		}

		if today.AddDate(0, 0, -3).Before(dueDate) {
			var adminUsers []*user.User
			err := r.db.Where("role = ?", "admin").Find(&adminUsers).Error
			if err != nil {
				return nil, err
			}
			for _, user := range adminUsers {
				reminder := fmt.Sprintf("Dear %s,This is a reminder that there is a task with ID %d that is due in less than 3 days on %s. Please take necessary actions to approve it.Best regards,Admin", user.Name, order.ID, order.DueDate)
				reminders = append(reminders, reminder)
			}
		}
	}
	result := map[string]interface{}{
		"rem": reminders,
	}

	return result, nil
}
