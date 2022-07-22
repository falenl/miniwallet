package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/falenl/miniwallet/usecase"
	mock_usecase "github.com/falenl/miniwallet/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestCreateAccountHandle(t *testing.T) {
	// c := NewCreateAccountHandler()
	ctx := context.Background()
	customerIDTemplate := uuid.New()
	token := "ygfyegf6ef87e6f6"
	tests := []struct {
		name       string
		mockFunc   func(mockAccountRepo *mock_usecase.MockAccountRepository, mockCustServ *mock_usecase.MockCustomerService, customerID uuid.UUID)
		expected   string
		customerID uuid.UUID
		isError    bool
	}{
		{
			name: "Success",
			mockFunc: func(mockAccountRepo *mock_usecase.MockAccountRepository, mockCustServ *mock_usecase.MockCustomerService, customerID uuid.UUID) {
				mockCustServ.EXPECT().Verify(ctx, customerID).Return(true)
				mockAccountRepo.EXPECT().Create(ctx, customerID).Return(token, nil)
			},
			customerID: customerIDTemplate,
			expected:   token,
		},
		{
			name: "error - nil cust id",
			mockFunc: func(mockAccountRepo *mock_usecase.MockAccountRepository, mockCustServ *mock_usecase.MockCustomerService, customerID uuid.UUID) {
			},
			expected: "",
			isError:  true,
		},
		{
			name: "error - customer not found",
			mockFunc: func(mockAccountRepo *mock_usecase.MockAccountRepository, mockCustServ *mock_usecase.MockCustomerService, customerID uuid.UUID) {
				mockCustServ.EXPECT().Verify(ctx, customerID).Return(false)
			},
			expected:   "",
			customerID: customerIDTemplate,
			isError:    true,
		},
		{
			name: "error - account repo error",
			mockFunc: func(mockAccountRepo *mock_usecase.MockAccountRepository, mockCustServ *mock_usecase.MockCustomerService, customerID uuid.UUID) {
				mockCustServ.EXPECT().Verify(ctx, customerID).Return(true)
				mockAccountRepo.EXPECT().Create(ctx, customerID).Return("", errors.New("database error"))
			},
			expected:   "",
			customerID: customerIDTemplate,
			isError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockAccountRepo := mock_usecase.NewMockAccountRepository(ctrl)
			mockCustServ := mock_usecase.NewMockCustomerService(ctrl)
			c := usecase.NewAccountService(mockAccountRepo, mockCustServ)

			test.mockFunc(mockAccountRepo, mockCustServ, test.customerID)
			tokenResult, err := c.Create(ctx, test.customerID)

			if test.expected != tokenResult {
				t.Errorf("Token Expected %s, but got %s", token, tokenResult)
			}

			if test.isError != (err != nil) {
				t.Errorf("Error Expected %s, but got %s", token, tokenResult)
			}
		})
	}
}
