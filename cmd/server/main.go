package main

import (
	"bankruptcy/internal/port"
	"bankruptcy/pkg/http"
	"fmt"
	"log"
)

func main() {
	fmt.Println("serving bankruptcy")

	httpError := make(chan error)
	go http.ServeHTTP(":8080", port.Get, port.Add, httpError)

	log.Fatal(<-httpError)
}
