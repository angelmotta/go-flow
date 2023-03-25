package main

import (
	"encoding/json"
	"fmt"
	"github.com/angelmotta/go-flow/internal/exchangestore"
	"log"
	"net/http"
)

type exchangeServer struct {
	store *exchangestore.ExchangeStore
}

func NewExchangeServer() *exchangeServer {
	store, err := exchangestore.New()
	if err != nil {
		log.Fatal("Error creating Database connection")
	}
	return &exchangeServer{store: store}
}

func (es *exchangeServer) getExchangeHandler(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprint(w, "20")
	res := struct {
		Status string `json:"status"`
	}{
		Status: "okkk",
	}

	js, err := json.Marshal(res)
	fmt.Println(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	fmt.Println("Golang with in-memory database")
	mux := http.NewServeMux()
	server := NewExchangeServer()
	mux.HandleFunc("/exchange", server.getExchangeHandler)
	http.ListenAndServe("localhost:5000", mux)
}
