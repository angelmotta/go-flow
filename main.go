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
	if req.URL.Path == "/exchange/" {
		// Exact match: request is plain "/exchange/", without trailing <sourceCurrency ID>
		if req.Method == http.MethodGet {
			es.getAllExchangesHandler(w, req)
		}
	} else {
		// Request has an Exchange ID, as in "/exchange/<idSrcCurrency-idDstCurrency>"
		// TODO: parse request and identify http method
		log.Printf("handling request with Exchange ID at %s\n", req.URL.Path)
	}
}

func (es *exchangeServer) getAllExchangesHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all exchanges available at %s\n", req.URL.Path)
	// TODO: implement method store.GetAllExchanges()
	// Idea: Get All keys and return a list based on a set
	response := struct {
		Currency []string `json:"exchangeList"`
	}{
		Currency: []string{"sol-dollar", "sol-euro", "dollar-sol", "dollar-euro"},
	}
	jsRes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsRes)
}

func main() {
	fmt.Println("Golang with in-memory database")
	mux := http.NewServeMux()
	server := NewExchangeServer()
	mux.HandleFunc("/exchange/", server.getExchangeHandler)

	log.Fatal(http.ListenAndServe("localhost:5000", mux))
}
