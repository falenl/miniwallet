package usecase

import (
	"context"

	"github.com/falenl/miniwallet/entity"
	"github.com/google/uuid"
)

//mockgen -source ./usecase/interfaces.go -destination ./usecase/mocks/interfaces_mock.go

type AccountRepository interface {
	Create(ctx context.Context, customerID uuid.UUID) (token string, err error)
	Authenticate(ctx context.Context, token string) (entity.Account, error)
}

type CustomerService interface {
	Verify(ctx context.Context, customerID uuid.UUID) bool
}

type WalletRepository interface {
	Enable(ctx context.Context, wallet *entity.Wallet) error
	GetByAccountID(ctx context.Context, accountID string) (entity.Wallet, error)
	Update(ctx context.Context, wallet *entity.Wallet) error
}

type TransactionRepository interface {
	Deposit(ctx context.Context, trx *entity.Transaction) error
	Withdraw(ctx context.Context, trx *entity.Transaction) error
}
