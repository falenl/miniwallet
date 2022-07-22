package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	Enabled  statusString = "enabled"
	Disabled statusString = "disabled"
)

type statusString string

type Wallet struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	OwnerID         uuid.UUID
	Status          statusString
	StatusUpdatedAt time.Time
	Balance         int64
}

func (w *Wallet) IsDisabled() bool {
	return w.Status == Disabled
}
