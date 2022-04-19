package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

const (
	DB      = "bky.json"
	DB_PATH = "/etc/bky/"
)

var (
	fullPath = fmt.Sprintf("%s%s", DB_PATH, DB)
)

type transaction struct {
	Date     time.Time `json:"date"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	To       string    `json:"to,omitempty"`
}

type group struct {
	Transactions []transaction `json:"transactions"`
}

func main() {
	var from string
	var to string
	flag.StringVar(&from, "from", "SEK", "Transaction currency")
	flag.StringVar(&to, "to", "EURO", "Destination currency")

	flag.Parse()

	if len(flag.Args()) > 0 {
		if in, err := strconv.ParseFloat(flag.Args()[0], 32); in > 0 {
			if err != nil {
				fmt.Printf("Cannot parse transaction amount %v\n", err)
			}
			amount := math.Round(in*100) / 100
			t := time.Now()
			err = createRepo()
			if err != nil {
				fmt.Println(err)
			}
			err = appendTransaction(&transaction{
				Date:     t,
				Amount:   amount,
				Currency: from,
				To:       to,
			})
			if err != nil {
				fmt.Println(err)
			}
		} else if err != nil {
			fmt.Println("Amount must be a number")
			return
		} else {
			fmt.Println("Amount must be greater than zero")
			return
		}
	} else {
		fmt.Println("No amount provided for transaction")
		return
	}
}

func createRepo() error {
	_, err := os.Stat(DB_PATH)
	if os.IsNotExist(err) {
		err = os.Mkdir(DB_PATH, 0755)
		return err
	}
	return nil
}

func appendTransaction(t *transaction) error {

	fs, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return err
	}
	size, _ := fs.Stat()
	buf := make([]byte, size.Size())
	fs.Read(buf)
	if len(buf) == 0 {
		first, err := json.Marshal(group{
			Transactions: []transaction{*t},
		})
		if err != nil {
			return err
		}
		fs.Write(first)
	} else {
		var g *group
		err := json.Unmarshal(buf, &g)

		if err != nil {
			return err
		}
		g.Transactions = append(g.Transactions, *t)
		b, err := json.Marshal(&g)
		if err != nil {
			return err
		}
		_, err = fs.WriteAt(b, 0)
		if err != nil {
			return err
		}
	}
	return nil
}
