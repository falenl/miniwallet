package usecase

import (
	"context"

	"github.com/google/uuid"
)

type AccountRepository interface {
	Create(ctx context.Context, customerID uuid.UUID) (token string, err error)
}

type CustomerService interface {
	Verify(ctx context.Context, customerID uuid.UUID) bool
}
