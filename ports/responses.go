package ports

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/falenl/miniwallet/errors"
	"github.com/google/uuid"
)

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
	StatusError   = "error"
)

var accountCtxKey = contextKey("Account ID")

type contextKey string
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type InitAccountResp struct {
	Token string `json:"token"`
}

type WalletResp struct {
	ID         uuid.UUID  `json:"id"`
	OwnedBy    uuid.UUID  `json:"owned_by"`
	Status     string     `json:"status"`
	EnabledAt  *time.Time `json:"enabled_at,omitempty"`
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	Balance    int64      `json:"balance"`
}

type DepositResp struct {
	Deposit DepositDetailResp `json:"deposit"`
}

type DepositDetailResp struct {
	ID          uuid.UUID `json:"id"`
	DepositedBy uuid.UUID `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      int64     `json:"amount"`
	RefID       uuid.UUID `json:"reference_id"`
}

type WithdrawalResp struct {
	WithDrawal WithdrawalDetailResp `json:"witdrawal"`
}

type WithdrawalDetailResp struct {
	ID          uuid.UUID `json:"id"`
	WithdrawnBy uuid.UUID `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      int64     `json:"amount"`
	RefID       uuid.UUID `json:"reference_id"`
}

func JSONResponse(w http.ResponseWriter, code int, output interface{}) error {
	response, err := json.Marshal(output)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err := jsonResponse(w, code, response); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func jsonResponse(w http.ResponseWriter, code int, output []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if _, err := w.Write(output); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func HandleErrReturn(w http.ResponseWriter, errInput *errors.Error) {
	status := StatusError
	if errInput.StatusCode == http.StatusBadRequest || errInput.StatusCode == http.StatusUnauthorized ||
		errInput.StatusCode == http.StatusOK {
		status = StatusFail
	}
	resp := Response{
		Status:  status,
		Message: errInput.Error(),
	}
	JSONResponse(w, errInput.StatusCode, resp)
	log.Printf("Error response: %s", errInput.Error())
}
