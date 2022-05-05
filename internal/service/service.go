package service

import (
	"bankruptcy/internal/core/adapter"
	"bankruptcy/internal/core/domain"
)

func RegisterTransactionFunc(r adapter.RegisterTransactionRepo) adapter.RegisterTransaction {
	return func(trs domain.Transaction) error {
		return r(trs)
	}
}

func ReadTransactionsFunc(r adapter.ReadTransactionsRepo) adapter.ReadTransactions {
	return func() (domain.Group, error) {
		return r()
	}
}
