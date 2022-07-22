package adapters

import (
	"context"
	"database/sql"
	"time"

	"github.com/falenl/miniwallet/entity"
)

const (
	getWalletByAccountIDQuery = `select id, account_id, owner_id, status, balance, status_updated_at from wallet where account_id = $1`
	createWalletQuery         = `insert into wallet (id, account_id, owner_id, status, balance, status_updated_at) 
								values($1, $2, $3, $4, $5, $6) 
								ON CONFLICT(id) 
								DO UPDATE SET status = excluded.status, status_updated_at = excluded.status_updated_at`
	updateWalletQuery = `update wallet set status = $1, status_updated_at = $2, balance = $3 where ID = $4`
)

type walletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *walletRepository {
	return &walletRepository{
		db,
	}
}

func (w *walletRepository) Enable(ctx context.Context, wallet *entity.Wallet) error {

	wallet.Status = entity.Enabled
	wallet.StatusUpdatedAt = time.Now()

	_, err := w.db.ExecContext(ctx, createWalletQuery, wallet.ID, wallet.AccountID, wallet.OwnerID,
		wallet.Status, wallet.Balance, wallet.StatusUpdatedAt.Format(time.RFC3339))

	if err != nil {
		return err
	}

	return nil
}

func (w *walletRepository) GetByAccountID(ctx context.Context, accountID string) (entity.Wallet, error) {
	row := w.db.QueryRowContext(ctx, getWalletByAccountIDQuery, accountID)
	if row.Err() != nil {
		return entity.Wallet{}, row.Err()
	}

	var wallet entity.Wallet
	var statusUpdatedAt string
	err := row.Scan(&wallet.ID, &wallet.AccountID, &wallet.OwnerID, &wallet.Status, &wallet.Balance, &statusUpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return entity.Wallet{}, err
	}

	if statusUpdatedAt != "" {
		wallet.StatusUpdatedAt, err = time.Parse(time.RFC3339, statusUpdatedAt)
		if err != nil {
			return entity.Wallet{}, err
		}
	}

	return wallet, nil
}

func (w *walletRepository) Update(ctx context.Context, wallet *entity.Wallet) error {

	_, err := w.db.ExecContext(ctx, updateWalletQuery, wallet.Status, wallet.StatusUpdatedAt.Format(time.RFC3339),
		wallet.Balance, wallet.ID)

	if err != nil {
		return err
	}

	return nil
}
