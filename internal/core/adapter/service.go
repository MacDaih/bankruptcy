package adapter

import "bankruptcy/internal/core/domain"

type RegisterTransaction func(domain.Transaction) error
type ReadTransactions func() (domain.Group, error)
