package ports

import (
	"context"
	"net/http"

	"github.com/falenl/miniwallet/errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Account interface {
	Create(ctx context.Context, customerID uuid.UUID) (string, error)
}

type accountHTTPHandler struct {
	account Account
}

func New(r *mux.Router, account Account) {
	a := accountHTTPHandler{
		account: account,
	}

	r.HandleFunc("/init", a.InitAccount)
}

func (a *accountHTTPHandler) InitAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := r.Context()

	customerID, err := uuid.Parse(r.FormValue("customer_xid"))
	if err != nil {
		errInput := errors.NewInvalidRequest("invalid customer_xid")
		HandleErrReturn(w, errInput)
		return
	}

	token, err := a.account.Create(ctx, customerID)
	if err != nil {
		HandleErrReturn(w, err.(*errors.Error))
		return
	}

	resp := Response{
		Status: StatusSuccess,
		Data: InitAccountResp{
			Token: token,
		},
	}
	JSONResponse(w, http.StatusCreated, resp)

}
