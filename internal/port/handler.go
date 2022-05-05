package port

import (
	"bankruptcy/internal/core/adapter"
	"bankruptcy/internal/core/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type handler struct {
	add adapter.RegisterTransaction
	get adapter.ReadTransactions
}

func NewHandler(
	register adapter.RegisterTransaction,
	read adapter.ReadTransactions,
) *handler {
	return &handler{
		add: register,
		get: read,
	}
}

func (h *handler) Get(w http.ResponseWriter, req *http.Request) {
	res, err := h.get()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *handler) Add(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var t domain.Transaction

	err = json.Unmarshal(b, &t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.add(t)

	fmt.Println(err)
}
