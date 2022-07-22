package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/falenl/miniwallet/adapters"
	"github.com/falenl/miniwallet/ports"
	"github.com/falenl/miniwallet/usecase"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	var (
		httpAddr = flag.String("http", ":80", "http listen address")
		errChan  = make(chan error)
	)

	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1").Subrouter()

	db, err := adapters.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to open database - error: %s", err.Error())
	}

	customerService := adapters.NewCustomerService()
	accountRepo := adapters.NewAccountRepository(db)
	account := usecase.NewAccountService(accountRepo, customerService)
	ports.NewAccountHandler(router, account)

	walletRepo := adapters.NewWalletRepository(db)
	wallet := usecase.NewWalletService(accountRepo, walletRepo)

	trxRepo := adapters.NewTransactionRepository(db)
	trx := usecase.NewTransactionService(wallet, trxRepo)
	ports.NewWalletHandler(router, wallet, trx)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("Server is up and listening on: ", *httpAddr)
		handler := cors.Default().Handler(router)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	err = <-errChan
	log.Println("Server is Stopped, - err: ", err)
}
