package main

import (
	"github.com/gorilla/mux"
	"net/http"

	as "github.com/my/repo/server/aerospike"
)

func main() {

	as.CheckDBConnection()

    r := mux.NewRouter()

    r.HandleFunc("/user/{id}", http.HandlerFunc(as.GetRecord))
    r.HandleFunc("/loaddata", http.HandlerFunc(as.PreloadCustomers))
    http.Handle("/", r)
	http.ListenAndServe(":8080", r)

}
