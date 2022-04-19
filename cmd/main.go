package main

import (
	"bankruptcy/internal/repository"
	"bankruptcy/internal/service"
	"flag"
	"fmt"
)

const (
	DB      = "bky.json"
	DB_PATH = "bankruptcy/"
)

func main() {

	var amount float64
	flag.Float64Var(&amount, "a", amount, "transaction amount")

	flag.Parse()

	trs := repository.NewTransactionRepo(DB_PATH, DB)

	srv := service.NewTransactionService(trs)

	if err := srv.Store(amount); err != nil {
		fmt.Println(err)
		return
	}
}
