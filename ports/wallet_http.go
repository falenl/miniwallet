package ports

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/falenl/miniwallet/entity"
	"github.com/falenl/miniwallet/errors"
	"github.com/gorilla/mux"
)

type Wallet interface {
	Authenticate(ctx context.Context, token string) (entity.Account, error)
	Enable(ctx context.Context, account entity.Account) (entity.Wallet, error)
	Get(ctx context.Context, accountID string) (entity.Wallet, error)
	Update(ctx context.Context, accountID string, editedWallet entity.Wallet) (entity.Wallet, error)
}

type walletHTTPHandler struct {
	wallet Wallet
}

func NewWalletHandler(r *mux.Router, wallet Wallet) {
	w := walletHTTPHandler{
		wallet: wallet,
	}

	r.HandleFunc("/wallet", w.Authenticate(w.Handler)).Methods(http.MethodPost, http.MethodGet, http.MethodPatch)
}

func (h *walletHTTPHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := ctx.Value(accountCtxKey)
	r.ParseForm()

	var (
		wallet entity.Wallet
		err    error
	)

	switch r.Method {
	case http.MethodPost:
		wallet, err = h.wallet.Enable(ctx, account.(entity.Account))
	case http.MethodGet:
		wallet, err = h.wallet.Get(ctx, account.(entity.Account).ID.String())
	case http.MethodPatch:
		isDisabledValue := r.FormValue("is_disabled")
		isDisabled, err2 := strconv.ParseBool(isDisabledValue)
		if err2 != nil {
			HandleErrReturn(w, err.(*errors.Error))
			return
		}
		wallet, err = h.patchWallet(ctx, account.(entity.Account).ID.String(), isDisabled)
	}

	if err != nil {
		HandleErrReturn(w, err.(*errors.Error))
		return
	}

	wr := WalletResp{
		ID:      wallet.ID,
		OwnedBy: wallet.OwnerID,
		Status:  string(wallet.Status),
		Balance: wallet.Balance,
	}
	if wallet.IsDisabled() {
		wr.DisabledAt = &wallet.StatusUpdatedAt
	} else {
		wr.EnabledAt = &wallet.StatusUpdatedAt
	}

	resp := Response{
		Status: StatusSuccess,
		Data:   wr,
	}

	JSONResponse(w, http.StatusCreated, resp)

}

func (h *walletHTTPHandler) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r.ParseForm()
		token := r.Header.Get("Authorization")

		tokens := strings.Split(token, " ")

		if len(tokens) != 2 {
			HandleErrReturn(w, errors.NewUnauthorized("missing authorization"))
			return
		}

		token = tokens[1]
		account, err := h.wallet.Authenticate(ctx, token)
		if err != nil {
			HandleErrReturn(w, errors.NewUnauthorized("missing authorization"))
			return
		}

		ctx = context.WithValue(ctx, accountCtxKey, account)
		req := r.WithContext(ctx)

		next(w, req)

	}
}

func (h *walletHTTPHandler) patchWallet(ctx context.Context, accountID string, isDisabled bool) (entity.Wallet, error) {
	wallet := entity.Wallet{
		Status: entity.Enabled,
	}
	if isDisabled {
		wallet.Status = entity.Disabled
	}

	wallet, err := h.wallet.Update(ctx, accountID, wallet)
	if err != nil {
		return entity.Wallet{}, err
	}

	return wallet, nil
}
