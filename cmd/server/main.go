package main

import (
	"bankruptcy/internal/port"
	"bankruptcy/pkg/http"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("serving bankruptcy")
	sig := make(chan os.Signal, 1)

	signal.Notify(sig)
	httpError := make(chan error)

	go http.ServeHTTP(":8080", port.Get, port.Add, httpError)

	select {
	case err := <-httpError:
		log.Fatalf("http server failed : %v", err)
	case <-sig:
		log.Println("http server shutting down")
	}
}
