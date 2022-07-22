package adapters

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/falenl/miniwallet/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {

	trx := entity.Transaction{

		Amount:    2000,
		RefID:     uuid.New(),
		Type:      entity.TypeDeposit,
		UpdatedBy: uuid.New(),
		Wallet: entity.Wallet{
			ID:      uuid.New(),
			Balance: 2000,
			Status:  entity.Enabled,
		},
	}
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(createTransactionQuery)).
			WithArgs(sqlmock.AnyArg(), trx.Wallet.ID, "success", trx.Amount, trx.RefID, trx.Type,
				sqlmock.AnyArg(), trx.UpdatedBy).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec(regexp.QuoteMeta(depositBalanceQuery)).
			WithArgs(trx.Amount, trx.Wallet.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		trxRepo := NewTransactionRepository(db)
		err = trxRepo.Deposit(ctx, &trx)

		assert.NoError(t, err)
	})
}

func TestWithdrawal(t *testing.T) {

	trx := entity.Transaction{

		Amount:    2000,
		RefID:     uuid.New(),
		Type:      entity.TypeWithdrawal,
		UpdatedBy: uuid.New(),
		Wallet: entity.Wallet{
			ID:      uuid.New(),
			Balance: 2000,
			Status:  entity.Enabled,
		},
	}
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(createTransactionQuery)).
			WithArgs(sqlmock.AnyArg(), trx.Wallet.ID, "success", trx.Amount, trx.RefID, trx.Type,
				sqlmock.AnyArg(), trx.UpdatedBy).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec(regexp.QuoteMeta(withdrawBalanceQuery)).
			WithArgs(trx.Amount, trx.Wallet.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		trxRepo := NewTransactionRepository(db)
		err = trxRepo.Withdraw(ctx, &trx)

		assert.NoError(t, err)
	})
}
