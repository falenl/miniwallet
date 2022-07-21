package adapters

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
