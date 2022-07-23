package adapters

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/falenl/miniwallet/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newWallet() entity.Wallet {
	return entity.Wallet{
		ID:              uuid.New(),
		AccountID:       uuid.New(),
		OwnerID:         uuid.New(),
		Status:          entity.Enabled,
		StatusUpdatedAt: time.Now(),
	}
}

func TestEnableWallet(t *testing.T) {
	ctx := context.Background()

	wallet := newWallet()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectExec(regexp.QuoteMeta(createWalletQuery)).
			WithArgs(wallet.ID, wallet.AccountID, wallet.OwnerID,
				wallet.Status, wallet.Balance, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

		walletRepo := NewWalletRepository(db)
		err = walletRepo.Enable(ctx, &wallet)
		assert.NoError(t, err)

	})
}

func TestGetAccountByID(t *testing.T) {
	ctx := context.Background()
	wallet := newWallet()
	accountID := "efefwffwf"

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery(regexp.QuoteMeta(getWalletByAccountIDQuery)).WithArgs(accountID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "account_id", "owner_id", "status", "balance", "status_updated_at"}).
				AddRow(wallet.ID, wallet.AccountID, wallet.OwnerID, wallet.Status, wallet.Balance, wallet.StatusUpdatedAt.Format(time.RFC3339)))

		walletRepo := NewWalletRepository(db)
		w, err := walletRepo.GetByAccountID(ctx, accountID)

		assert.Equal(t, wallet.ID, w.ID)
		assert.NoError(t, err)
	})
}

func TestUpdateWallet(t *testing.T) {
	ctx := context.Background()

	wallet := newWallet()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectExec(regexp.QuoteMeta(updateWalletQuery)).
			WithArgs(wallet.Status, sqlmock.AnyArg(), wallet.Balance, wallet.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		walletRepo := NewWalletRepository(db)
		err = walletRepo.Update(ctx, &wallet)
		assert.NoError(t, err)

	})
}
