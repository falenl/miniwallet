package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/falenl/miniwallet/entity"
	errs "github.com/falenl/miniwallet/errors"

	"github.com/google/uuid"
)

type walletService struct {
	accountRepo AccountRepository
	walletRepo  WalletRepository
}

//NewWalletService returns walletService struct
//accepts AccountRepository and WalletRepository as parameters
func NewWalletService(accountRepo AccountRepository, walletRepo WalletRepository) *walletService {
	return &walletService{
		accountRepo: accountRepo,
		walletRepo:  walletRepo,
	}
}

//Enable will create a wallet or enable a disabled wallet and return error if any
func (w *walletService) Enable(ctx context.Context, account entity.Account) (entity.Wallet, error) {
	wallet, err := w.walletRepo.GetByAccountID(ctx, account.ID.String())
	if err != nil {
		return entity.Wallet{}, errs.NewInternalServer(err.Error())
	}

	//create new ID if not exists
	if wallet.ID == uuid.Nil {

		wallet.ID = uuid.New()
		wallet.AccountID = account.ID
		wallet.OwnerID = account.CustomerID

	} else if !wallet.IsDisabled() {

		return entity.Wallet{}, errs.NewExpected("wallet has been enabled")
	}

	if err = w.walletRepo.Enable(ctx, &wallet); err != nil {
		return entity.Wallet{}, errs.NewInternalServer(err.Error())
	}

	return wallet, nil
}

//GetByAccountID returns wallet by account ID
func (w *walletService) Get(ctx context.Context, accountID string) (entity.Wallet, error) {
	wallet, err := w.walletRepo.GetByAccountID(ctx, accountID)
	if err != nil {
		return entity.Wallet{}, errs.NewInternalServer(err.Error())
	}

	if wallet.IsDisabled() {
		return entity.Wallet{}, errs.NewExpected("account has been disabled")
	}

	return wallet, nil
}

func (w *walletService) Authenticate(ctx context.Context, token string) (entity.Account, error) {
	account, err := w.accountRepo.Authenticate(ctx, token)
	errUnauthorized := errs.NewUnauthorized("failed to authenticate")
	if err != nil {
		err = fmt.Errorf("failed to authenticate - %w", err)
		log.Println(err.Error())
		return entity.Account{}, errUnauthorized
	}

	return account, nil
}

//Update set the status of a wallet
func (w *walletService) Update(ctx context.Context, accountID string, editedWallet entity.Wallet) (entity.Wallet, error) {
	oriWallet, err := w.Get(ctx, accountID)
	if err != nil {
		return entity.Wallet{}, err
	}

	editedWallet.ID = oriWallet.ID
	editedWallet.OwnerID = oriWallet.OwnerID
	editedWallet.Balance = oriWallet.Balance
	if oriWallet.Status != editedWallet.Status {
		editedWallet.StatusUpdatedAt = time.Now()
	}

	err = w.walletRepo.Update(ctx, &editedWallet)
	if err != nil {
		return entity.Wallet{}, err
	}

	return editedWallet, nil
}
