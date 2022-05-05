package adapter

import "bankruptcy/internal/core/domain"

type RegisterTransactionRepo func(domain.Transaction) error
type ReadTransactionsRepo func() (domain.Group, error)
