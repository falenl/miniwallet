package adapters

import (
	"context"
	"database/sql"
	"time"

	"github.com/falenl/miniwallet/entity"
	"github.com/google/uuid"
)

const (
	createTransactionQuery = `INSERT INTO transactions(id, wallet_id, status, amount, reference_id, transaction_type, updated_at, updated_by) 
							values ($1, $2, $3, $4, $5, $6, $7, $8)`
	depositBalanceQuery  = `UPDATE wallet SET balance = balance + $1 where ID = $2`
	withdrawBalanceQuery = `UPDATE wallet SET balance = balance - $1 where ID = $2`
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{
		db,
	}
}

func (t *transactionRepository) Deposit(ctx context.Context, trx *entity.Transaction) error {
	trx.Type = entity.TypeDeposit

	return t.createTransaction(ctx, trx, depositBalanceQuery)
}

func (t *transactionRepository) Withdraw(ctx context.Context, trx *entity.Transaction) error {
	trx.Type = entity.TypeWithdrawal

	return t.createTransaction(ctx, trx, withdrawBalanceQuery)
}

func (t *transactionRepository) createTransaction(ctx context.Context, trx *entity.Transaction, updateQuery string) (err error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	trx.ID = uuid.New()

	_, err = tx.ExecContext(ctx, createTransactionQuery, trx.ID, trx.Wallet.ID, "success", trx.Amount, trx.RefID, trx.Type,
		trx.UpdatedAt.Format(time.RFC3339), trx.UpdatedBy)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, updateQuery, trx.Amount, trx.Wallet.ID)

	if err != nil {
		return err
	}

	return nil
}
