package adapters

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/falenl/miniwallet/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	customerID := uuid.New()
	ctx := context.Background()
	token := "yugewf873ufeh9"
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery(regexp.QuoteMeta(createAccountQuery)).WithArgs(sqlmock.AnyArg(), customerID, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"token"}).AddRow(token))

		accRepo := NewAccountRepository(db)
		tokenResult, err2 := accRepo.Create(ctx, customerID)
		assert.Equal(t, token, tokenResult)
		assert.NoError(t, err2)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		errDatabase := errors.New("err database")
		mock.ExpectQuery(regexp.QuoteMeta(createAccountQuery)).WithArgs(sqlmock.AnyArg(), customerID, sqlmock.AnyArg()).WillReturnError(errDatabase)

		accRepo := NewAccountRepository(db)
		tokenResult, err2 := accRepo.Create(ctx, customerID)
		assert.Empty(t, tokenResult)
		assert.ErrorIs(t, errDatabase, err2)
	})
}

func TestAuthenticate(t *testing.T) {
	ctx := context.Background()
	token := "yugewf873ufeh9"
	account := entity.Account{
		ID:         uuid.New(),
		CustomerID: uuid.New(),
	}
	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery(regexp.QuoteMeta(getAccountQuery)).WithArgs(token).
			WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "token"}).
				AddRow(account.ID.String(), account.CustomerID.String(), token))

		accRepo := NewAccountRepository(db)
		accResult, err2 := accRepo.Authenticate(ctx, token)
		assert.Equal(t, account, accResult)
		assert.NoError(t, err2)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery(regexp.QuoteMeta(getAccountQuery)).WithArgs(token).
			WillReturnError(sql.ErrNoRows)

		accRepo := NewAccountRepository(db)
		_, err2 := accRepo.Authenticate(ctx, token)

		assert.Error(t, err2)
	})

}
