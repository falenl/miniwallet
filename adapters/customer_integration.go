package adapters

import (
	"context"

	"github.com/google/uuid"
)

type customerService struct {
}

//Verify will always return true for poc.
//customer verification are done in a separate customer service
func (c *customerService) Verify(ctx context.Context, customerID uuid.UUID) bool {
	return true
}

func NewCustomerService() *customerService {
	return &customerService{}
}
