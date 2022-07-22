package usecase

import (
	"context"
	"time"

	"github.com/falenl/miniwallet/entity"
	errs "github.com/falenl/miniwallet/errors"
	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
)

type transactionService struct {
	ws      *walletService
	trxRepo TransactionRepository
}

func NewTransactionService(ws *walletService, trxRepo TransactionRepository) *transactionService {
	return &transactionService{
		ws:      ws,
		trxRepo: trxRepo,
	}
}

func (ts *transactionService) createTransaction(ctx context.Context, account entity.Account, refID string, amount int64) (entity.Transaction, error) {
	var (
		trx entity.Transaction
		err error
	)

	//input validation
	trx.Amount = amount
	trx.RefID, err = uuid.Parse(refID)
	if err != nil {
		return trx, errs.NewInvalidRequest("reference ID must be a UUID")
	}

	if amount <= 0 {
		return trx, errs.NewInvalidRequest("amount must be greater than 0")
	}

	wallet, err := ts.ws.Get(ctx, account.ID.String())
	if err != nil {
		return trx, errs.NewInternalServer(err.Error())
	}

	trx.UpdatedAt = time.Now()
	trx.UpdatedBy = account.CustomerID
	trx.Wallet = wallet

	return trx, nil
}

func (ts *transactionService) Deposit(ctx context.Context, account entity.Account, refID string, amount int64) (entity.Transaction, error) {
	trx, err := ts.createTransaction(ctx, account, refID, amount)
	if err != nil {
		return trx, err
	}

	err = ts.trxRepo.Deposit(ctx, &trx)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return entity.Transaction{}, errs.NewExpected("reference_id is already exists")
			}
		}
		return trx, errs.NewInternalServer(err.Error())
	}

	return trx, nil
}

func (ts *transactionService) Withdraw(ctx context.Context, account entity.Account, refID string, amount int64) (entity.Transaction, error) {
	trx, err := ts.createTransaction(ctx, account, refID, amount)
	if err != nil {
		return trx, err
	}

	err = ts.trxRepo.Withdraw(ctx, &trx)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return entity.Transaction{}, errs.NewExpected("reference_id is already exists")
			}
		}
		return entity.Transaction{}, errs.NewInternalServer(err.Error())
	}

	return trx, nil
}
