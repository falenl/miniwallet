package entity

import (
	"time"

	"github.com/google/uuid"
)

type transactionType string

const TypeDeposit transactionType = "deposit"
const TypeWithdrawal transactionType = "withdrawal"

type Transaction struct {
	ID        uuid.UUID
	Amount    int64
	RefID     uuid.UUID
	Wallet    Wallet
	Type      transactionType
	UpdatedBy uuid.UUID
	UpdatedAt time.Time
}
