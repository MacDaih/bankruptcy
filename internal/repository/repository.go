package repository

import (
	"bankruptcy/internal/core/adapter"
	"bankruptcy/internal/core/domain"
	"encoding/json"
	"os"
)

func RegisterTransactionFunc(fs *os.File) adapter.RegisterTransactionRepo {
	return func(trs domain.Transaction) error {
		size, _ := fs.Stat()
		buf := make([]byte, size.Size())
		fs.Read(buf)
		if len(buf) == 0 {
			first, err := json.Marshal(domain.Group{
				Transactions: []domain.Transaction{trs},
			})
			if err != nil {
				return err
			}
			_, err = fs.Write(first)
			return err
		}

		var g *domain.Group
		err := json.Unmarshal(buf, &g)

		if err != nil {
			return err
		}
		g.Transactions = append(g.Transactions, trs)
		b, err := json.Marshal(&g)
		if err != nil {
			return err
		}
		_, err = fs.WriteAt(b, 0)
		if err != nil {
			return err
		}
		return nil
	}
}

func ReadTransactionsFunc(fs *os.File) adapter.ReadTransactionsRepo {
	return func() (domain.Group, error) {
		size, _ := fs.Stat()

		buf := make([]byte, size.Size())

		fs.Read(buf)

		var g domain.Group
		err := json.Unmarshal(buf, &g)

		return g, err
	}
}
