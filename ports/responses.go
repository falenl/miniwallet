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

type contextKey string

var accountCtxKey = contextKey("Account ID")

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
