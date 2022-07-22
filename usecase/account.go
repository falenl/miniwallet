package usecase

import (
	"context"

	errs "github.com/falenl/miniwallet/errors"
	"github.com/google/uuid"
)

type accountService struct {
	accountRepo     AccountRepository
	customerService CustomerService
}

//NewAccountService returns accountService to create account usecase
//accepts AccountRepository and CustomerService as parameters
func NewAccountService(accountRepo AccountRepository, customerService CustomerService) *accountService {
	return &accountService{
		accountRepo:     accountRepo,
		customerService: customerService,
	}
}

//Create returns a token string and error if any.
func (c *accountService) Create(ctx context.Context, customerID uuid.UUID) (string, error) {
	if customerID == uuid.Nil {
		return "", errs.NewInvalidRequest("invalid customer_xid")
	}
	if !c.verifyCustomer(ctx, customerID) {
		return "", errs.NewExpected("customer not found")
	}

	token, err := c.accountRepo.Create(ctx, customerID)
	if err != nil {
		return "", errs.NewInternalServer(err.Error())
	}
	return token, nil
}

func (c *accountService) verifyCustomer(ctx context.Context, customerID uuid.UUID) bool {
	return c.customerService.Verify(ctx, customerID)
}
