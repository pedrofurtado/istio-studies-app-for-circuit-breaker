package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", Run)
	http.ListenAndServe(":8080", nil)
}

func Run(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("FORCE_ERROR") == "yes" {
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		w.WriteHeader(http.StatusGatewayTimeout)
    w.Write([]byte("Istio studies app for circuit breaker - FAIL"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Istio studies app for circuit breaker - OK"))
}
