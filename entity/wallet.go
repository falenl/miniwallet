package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

type Wallet struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	OwnerID         uuid.UUID
	Status          string
	StatusUpdatedAt time.Time
	Balance         int64
}

type Transaction struct {
	ID        uuid.UUID
	Amount    int64
	RefID     uuid.UUID
	Wallet    Wallet
	UpdatedBy uuid.UUID
	UpdatedAt time.Time
}

func (w *Wallet) IsDisabled() bool {
	return w.Status == DISABLED
}
