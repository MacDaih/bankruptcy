package service

import (
	"bankruptcy/internal/core"
	"bankruptcy/internal/repository"
	"time"
)

type TransactionService struct {
	Repository *repository.TransactionRepo
}

func NewTransactionService(repository *repository.TransactionRepo) TransactionService {
	return TransactionService{
		Repository: repository,
	}
}

func (srv *TransactionService) Store(in float64) error {
	return srv.Repository.AppendTransaction(core.Transaction{
		Date:   time.Now(),
		Amount: in,
	})
}
