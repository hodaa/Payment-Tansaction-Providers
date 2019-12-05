package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/payment/transaction", get).Methods(http.MethodGet)

	return r

}
