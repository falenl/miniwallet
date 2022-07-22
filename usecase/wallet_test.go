package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/falenl/miniwallet/entity"
	"github.com/falenl/miniwallet/usecase"
	mock_usecase "github.com/falenl/miniwallet/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

type dependency struct {
	accRepo     *mock_usecase.MockAccountRepository
	custService *mock_usecase.MockCustomerService
	walletRepo  *mock_usecase.MockWalletRepository
}

func newDependency(t *testing.T) *dependency {
	ctrl := gomock.NewController(t)
	accRepo := mock_usecase.NewMockAccountRepository(ctrl)
	custService := mock_usecase.NewMockCustomerService(ctrl)
	walletRepo := mock_usecase.NewMockWalletRepository(ctrl)

	return &dependency{
		accRepo:     accRepo,
		custService: custService,
		walletRepo:  walletRepo,
	}
}

func TestEnableWallet(t *testing.T) {
	ctx := context.Background()
	account := entity.Account{
		ID:         uuid.New(),
		CustomerID: uuid.New(),
	}
	wallet := entity.Wallet{
		ID:              uuid.New(),
		AccountID:       account.ID,
		OwnerID:         account.CustomerID,
		Status:          entity.Disabled,
		StatusUpdatedAt: time.Now(),
	}

	tests := []struct {
		name     string
		mockFunc func(d *dependency)
		expected entity.Wallet
		isError  bool
	}{
		{
			name: "Success",
			mockFunc: func(d *dependency) {
				d.walletRepo.EXPECT().GetByAccountID(ctx, account.ID.String()).Return(wallet, nil)
				d.walletRepo.EXPECT().Enable(ctx, &wallet).Return(nil)
			},
			expected: wallet,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := newDependency(t)
			w := usecase.NewWalletService(d.accRepo, d.walletRepo)

			test.mockFunc(d)
			walletResult, err := w.Enable(ctx, account)

			if test.expected != walletResult {
				t.Errorf("result Expected %v, but got %v", test.expected, walletResult)
			}

			if test.isError != (err != nil) {
				t.Errorf("error Expected is different")
			}
		})
	}
}
