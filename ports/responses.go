package ports

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/falenl/miniwallet/errors"
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
	if errInput.StatusCode == http.StatusBadRequest {
		status = StatusFail
	}
	resp := Response{
		Status:  status,
		Message: errInput.Error(),
	}
	JSONResponse(w, errInput.StatusCode, resp)
	log.Printf("Error response: %s", errInput.Error())
}
