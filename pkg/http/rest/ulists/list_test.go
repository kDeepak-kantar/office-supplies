package ulists

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"net/http"
	"net/http/httptest"
	"net/url"

	ulistdomainmock "github.com/Deepak/pkg/domain/ulist/mocks"
	"github.com/Deepak/pkg/storage/db/userlist"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestOrderRequest struct {
	InputRequest  userlist.Order
	Error         string
	RespBody      string
	RespBodyOrder userlist.Order
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}

func MockJsonPost(c *gin.Context, content interface{}, params gin.Params) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)
	c.Params = params

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func MockJsonDelete(c *gin.Context, params gin.Params) {
	c.Request.Method = "DELETE"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)
	c.Params = params
}

// func TestCreateUserList(t *testing.T) {
// 	dmock := ulistdomainmock.NewDomain(t)
// 	officesupplies, _ := Init(Input{
// 		Ulist: dmock,
// 	})
// 	params := []gin.Param{
// 		{
// 			Key:   "id",
// 			Value: "1",
// 		},
// 	}
// 	cases := map[int]TestOrderRequest{
// 		1: {
// 			InputRequest: userlist.Order{
// 				UserID: nil,
// 				Items: []userlist.Item{
// 					{
// 						ItemID:   "3",
// 						Quantity: "2",
// 						OrderID:  "6",
// 					},
// 					{
// 						ItemID:   "5",
// 						Quantity: "4",
// 						OrderID:  "6",
// 					},
// 				},
// 				EmpName:       "Deepak",
// 				EmpEmail:      "Deepak@blackwoodseven.com",
// 				RequestedDate: "15/02/2023",
// 				DueDate:       "15/02/2023",
// 				Status:        "pending",
// 			},
// 			Error:    "internal server error",
// 			RespBody: "order not placed",
// 		},
// 		2: {
// 			InputRequest: userlist.Order{
// 				UserID: nil,
// 				Items: []userlist.Item{
// 					{
// 						ItemID:   "3",
// 						Quantity: "2",
// 						OrderID:  "6",
// 					},
// 					{
// 						ItemID:   "5",
// 						Quantity: "4",
// 						OrderID:  "6",
// 					},
// 				},
// 				EmpName:       "Deepak",
// 				EmpEmail:      "Deepak@blackwoodseven.com",
// 				RequestedDate: "15/02/2023",
// 				DueDate:       "15/02/2023",
// 				Status:        "pending",
// 			},
// 			Error:    "",
// 			RespBody: "order placed sucessfully",
// 		},
// 	}
// 	for i, v := range cases {
// 		w := httptest.NewRecorder()
// 		c := GetTestGinContext(w)
// 		MockJsonPost(c, v.InputRequest, params)
// 		suppliesreq := &userlist.Order{
// 			UserID:        v.InputRequest.UserID,
// 			Items:         v.InputRequest.Items,
// 			EmpName:       v.InputRequest.EmpName,
// 			EmpEmail:      v.InputRequest.EmpEmail,
// 			RequestedDate: v.InputRequest.RequestedDate,
// 			DueDate:       v.InputRequest.DueDate,
// 			Status:        v.InputRequest.Status,
// 		}
// 		switch i {
// 		case 1:
// 			dmock.On("CreateUserList", suppliesreq).Return(fmt.Errorf("internal server error")).Once()
// 		case 2:
// 			dmock.On("CreateUserList", suppliesreq).Return(nil).Once()
// 		}
// 		officesupplies.CreateUserList(c)

// if v.Error == "" {
// 	assert.EqualValues(t, http.StatusOK, w.Code)
// } else {
// 	assert.EqualValues(t, http.StatusUnprocessableEntity, w.Code)
// }
// 	}

// }
func TestCreateUserList(t *testing.T) {
	dmock := ulistdomainmock.NewDomain(t)
	officesupplies, _ := Init(Input{
		Ulist: dmock,
	})
	params := []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}
	cases := map[int]TestOrderRequest{
		1: {
			InputRequest: userlist.Order{
				UserID: nil,
				Items: []userlist.Item{
					{
						ItemID:   "3",
						Quantity: "2",
						OrderID:  "6",
					},
					{
						ItemID:   "5",
						Quantity: "4",
						OrderID:  "6",
					},
				},
				EmpName:       "Deepak",
				EmpEmail:      "Deepak@blackwoodseven.com",
				RequestedDate: "15/02/2023",
				DueDate:       "15/02/2023",
				Status:        "pending",
			},
			Error:    "internal server",
			RespBody: "order not placed",
		},
		2: {
			InputRequest: userlist.Order{
				UserID: nil,
				Items: []userlist.Item{
					{
						ItemID:   "3",
						Quantity: "2",
						OrderID:  "6",
					},
					{
						ItemID:   "5",
						Quantity: "4",
						OrderID:  "6",
					},
				},
				EmpName:       "Deepak",
				EmpEmail:      "Deepak@blackwoodseven.com",
				RequestedDate: "15/02/2023",
				DueDate:       "15/02/2023",
				Status:        "pending",
			},
			Error:    "",
			RespBody: "order placed sucessfully",
		},
	}
	for i, v := range cases {
		w := httptest.NewRecorder()
		c := GetTestGinContext(w)
		MockJsonPost(c, v.InputRequest, params)
		suppliesreq := &userlist.Order{
			UserID:        v.InputRequest.UserID,
			Items:         v.InputRequest.Items,
			EmpName:       v.InputRequest.EmpName,
			EmpEmail:      v.InputRequest.EmpEmail,
			RequestedDate: v.InputRequest.RequestedDate,
			DueDate:       v.InputRequest.DueDate,
			Status:        v.InputRequest.Status,
		}
		switch i {
		case 1:
			dmock.On("CreateUserList", suppliesreq).Return(fmt.Errorf("internal server error")).Once()
		case 2:
			dmock.On("CreateUserList", suppliesreq).Return(nil).Once()
		}
		officesupplies.CreateUserList(c)
		if v.Error == "" {
			assert.EqualValues(t, http.StatusOK, w.Code)
		} else {
			assert.EqualValues(t, http.StatusUnprocessableEntity, w.Code)
		}
	}
}
