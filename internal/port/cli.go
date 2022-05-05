package port

import (
	"bankruptcy/internal/core/adapter"
	"bankruptcy/internal/core/domain"
	"fmt"
	"time"
)

type cli struct {
	add adapter.RegisterTransaction
	get adapter.ReadTransactions
}

func NewCli(
	register adapter.RegisterTransaction,
	read adapter.ReadTransactions,
) *cli {
	return &cli{
		add: register,
		get: read,
	}
}

func (c *cli) Add(input float64) {
	if input <= 0 {
		fmt.Println("invalid amount")
		return
	}
	err := c.add(domain.Transaction{
		Date:   time.Now(),
		Amount: input,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (c *cli) ReadAll() {
	res, err := c.get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
