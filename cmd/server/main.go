package main

import (
	"bankruptcy/internal/port"
	"bankruptcy/internal/repository"
	"bankruptcy/internal/service"
	"bankruptcy/pkg/http"
	"bankruptcy/pkg/store"
	"fmt"
	"log"
	"os"
	"os/signal"
)

const (
	DB      = "bky.json"
	DB_PATH = "bankruptcy/"
)

func main() {
	fmt.Println("serving bankruptcy")
	sig := make(chan os.Signal, 1)

	signal.Notify(sig)
	httpError := make(chan error)

	f, err := store.GetFile(DB_PATH, DB)

	if err != nil {
		fmt.Println("not enough argument")
		return
	}
	defer f.Close()

	handler := port.NewHandler(
		service.RegisterTransactionFunc(repository.RegisterTransactionFunc(f)),
		service.ReadTransactionsFunc(repository.ReadTransactionsFunc(f)),
	)
	go http.ServeHTTP(":8080", handler.Add, handler.Get, httpError)

	select {
	case err := <-httpError:
		log.Fatalf("http server failed : %v", err)
	case <-sig:
		log.Println("http server shutting down")
	}
}
