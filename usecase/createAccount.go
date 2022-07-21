package usecase

import (
	"context"

	"github.com/falenl/miniwallet/error"
	"github.com/google/uuid"
)

type AccountRepository interface {
}

type createAccountHandler struct {
	accountRepo AccountRepository
}

func NewCreateAccountHandler(accountRepo AccountRepository) createAccountHandler {
	return createAccountHandler{
		accountRepo: accountRepo,
	}
}

func (c createAccountHandler) Handle(ctx context.Context, customerID uuid.UUID) (string, error) {
	if customerID == uuid.Nil {
		return "", error.NewErrInvalidRequest("invalid customer_xid")
	}
}
