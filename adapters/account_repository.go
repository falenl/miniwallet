package adapters

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"

	"github.com/falenl/miniwallet/entity"
	"github.com/google/uuid"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *accountRepository {
	return &accountRepository{
		db,
	}
}

const (
	createAccountQuery = `insert into account(id, customer_id, token) values ($1, $2, $3) 
							ON CONFLICT (customer_id) DO UPDATE SET customer_id = excluded.customer_id 
							RETURNING token`

	getAccountQuery = `select id, customer_id, token
						from account
						where token = $1`
)

func (a *accountRepository) Create(ctx context.Context, customerID uuid.UUID) (string, error) {
	token := generateSecureToken(25)

	row := a.db.QueryRowContext(ctx, createAccountQuery, uuid.New().String(), customerID, token)
	if row.Err() != nil {
		return "", row.Err()
	}

	row.Scan(&token)

	return token, nil
}

func (a *accountRepository) Authenticate(ctx context.Context, token string) (entity.Account, error) {

	row := a.db.QueryRowContext(ctx, getAccountQuery, token)
	if row.Err() != nil {
		return entity.Account{}, row.Err()
	}

	var account entity.Account

	err := row.Scan(&account.ID, &account.CustomerID, &token)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func generateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
