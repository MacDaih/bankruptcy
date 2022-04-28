package main

import (
	"bankruptcy/internal/port"
	"bankruptcy/internal/repository"
	"bankruptcy/internal/service"
	"flag"
	"fmt"
)

const (
	DB      = "bky.json"
	DB_PATH = "bankruptcy/"

	GET = "get"
	ADD = "add"
)

func main() {
	var amount float64
	flag.Float64Var(&amount, "a", amount, "transaction amount")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("not enough argument")
		return
	}
	cmd := flag.Args()

	cli := port.NewCli(
		service.RegisterTransactionFunc(repository.RegisterTransactionFunc(DB_PATH, DB)),
		service.ReadTransactionsFunc(repository.ReadTransactionsFunc(DB_PATH, DB)),
	)
	switch cmd[0] {
	case GET:
		cli.ReadAll()
	case ADD:
		cli.Add(amount)
	default:
		fmt.Println("unknown command, type --help for command details")
	}
}
