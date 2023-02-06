package ulist

import (
	"github.com/Deepak/pkg/storage/userlist"
)

func (d *domain) CreateUserList(c *userlist.Order) error {

	err := d.UserList.CreateUserList(c)

	if err != nil {
		return nil
	}
	return nil
}
func (r *domain) SendRemainder() (map[string]interface{}, error) {
	return r.UserList.SendRemainder()
}
func (r *domain) GetAllUserLists() ([]*userlist.Order, error) {
	return r.UserList.GetAllist()
}
func (r *domain) GetUserLists(id string) ([]*userlist.Order, error) {
	return r.UserList.GetUserList(id)
}

func (r *domain) GetAllApprovedUserLists() ([]*userlist.Order, error) {
	return r.UserList.GetAllApproved()
}
func (r *domain) GetAllNotApprovedUserLists() ([]*userlist.Order, error) {
	return r.UserList.GetAllNotApproved()
}
func (r *domain) GetUserListStatus(id int) (string, error) {
	userDetails, err := r.UserList.GetOrderByStringiId(id)
	if userDetails == nil || err != nil {
		return "", err
	}
	return userDetails.Status, nil
}
func (r *domain) UpdateUserListstat(id int, status string) (*userlist.Order, error) {
	statusdetails, err := r.UserList.GetOrderByStringiId(id)
	if statusdetails == nil || err != nil {
		return nil, ErrInvalidInputlist
	}
	statusdetails.Status = status
	err = r.UserList.UpdateList(statusdetails)
	if err != nil {
		return nil, err
	}
	return statusdetails, nil
}

// func (r *domain) UpdateUserList(c *userlist.Order) (*userlist.Order, error) {
// 	statusdetails, err := r.UserList.GetOrderByStringiId(int(c.ID))
// 	if statusdetails == nil || err != nil {
// 		return nil, ErrInvalidInputlist
// 	}
// 	statusdetails.Items = c.Items
// 	err = r.UserList.UpdateList(statusdetails)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return statusdetails, nil
// }
