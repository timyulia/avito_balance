package handler

import (
	"balance"
	"balance/pkg/service"
	mock_service "balance/pkg/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_addMoney(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBilling, acc balance.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            balance.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id": 1, "amount": 10, "reason": "no reason"}`,
			inputUser: balance.User{
				Id:     1,
				Amount: 10,
				Reason: "no reason",
			},
			mockBehavior: func(r *mock_service.MockBilling, acc balance.User) {
				r.EXPECT().AddMoney(acc).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"ok"}`,
		},
		{
			name:                 "bad request",
			inputBody:            `{"id": "boy", "amount": 10, "reason": "no reason"}`,
			inputUser:            balance.User{},
			mockBehavior:         func(r *mock_service.MockBilling, acc balance.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"json: cannot unmarshal string into Go struct field User.id of type int"}`,
		},
		{
			name:      "service error",
			inputBody: `{"id": 1, "amount": 10, "reason": "no reason"}`,
			inputUser: balance.User{
				Id:     1,
				Amount: 10,
				Reason: "no reason",
			},
			mockBehavior: func(r *mock_service.MockBilling, acc balance.User) {
				r.EXPECT().AddMoney(acc).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBilling(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Billing: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/add", handler.addMoney)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_reserve(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBilling, ord balance.Order)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            balance.Order
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"user_id": 14, "amount": 10, "order_id": 5, "service_id": 6}`,
			inputUser: balance.Order{
				UserId:    14,
				ServiceId: 6,
				OrderId:   5,
				Amount:    10,
			},
			mockBehavior: func(r *mock_service.MockBilling, ord balance.Order) {
				r.EXPECT().Reserve(ord).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"ok"}`,
		},
		{
			name:                 "ok",
			inputBody:            `{"user_id": "boy", "amount": 10, "order_id": 5, "service_id": 6}`,
			inputUser:            balance.Order{},
			mockBehavior:         func(r *mock_service.MockBilling, ord balance.Order) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"json: cannot unmarshal string into Go struct field Order.user_id of type int"}`,
		},
		{
			name:      "ok",
			inputBody: `{"user_id": 14, "amount": 10, "order_id": 5, "service_id": 6}`,
			inputUser: balance.Order{
				UserId:    14,
				ServiceId: 6,
				OrderId:   5,
				Amount:    10,
			},
			mockBehavior: func(r *mock_service.MockBilling, ord balance.Order) {
				r.EXPECT().Reserve(ord).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBilling(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Billing: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/reserve", handler.reserve)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/reserve",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
