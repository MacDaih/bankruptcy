package main

import (
	"bankruptcy/internal/port"
	"bankruptcy/internal/repository"
	"bankruptcy/internal/service"
	"bankruptcy/pkg/store"
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

	f, err := store.GetFile(DB_PATH, DB)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	cli := port.NewCli(
		service.RegisterTransactionFunc(repository.RegisterTransactionFunc(f)),
		service.ReadTransactionsFunc(repository.ReadTransactionsFunc(f)),
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
