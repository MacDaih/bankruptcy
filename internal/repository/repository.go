package repository

import (
	"bankruptcy/internal/core"
	"encoding/json"
	"os"
)

type TransactionRepo struct {
	path string
	file string
}

func NewTransactionRepo(path string, file string) *TransactionRepo {
	return &TransactionRepo{
		path: path,
		file: file,
	}
}

func createStore(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		return err
	}
	return nil
}

func (t *TransactionRepo) AppendTransaction(trs core.Transaction) error {

	if err := createStore(t.path); err != nil {
		return err
	}

	fs, err := os.OpenFile(t.path+t.file, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return err
	}

	size, _ := fs.Stat()
	buf := make([]byte, size.Size())
	fs.Read(buf)
	if len(buf) == 0 {
		first, err := json.Marshal(core.Group{
			Transactions: []core.Transaction{trs},
		})
		if err != nil {
			return err
		}
		_, err = fs.Write(first)
		return err
	}

	var g *core.Group
	err = json.Unmarshal(buf, &g)

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
