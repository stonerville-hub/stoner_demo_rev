package main

import (
	"net/http"

	"github.com/gorilla/mux"

	as "github.com/my/repo/server/aerospike"
	util "github.com/my/repo/utility"
)

func main() {
	as.CheckDBConnection()

    r := mux.NewRouter()
    r.HandleFunc(util.HOME, http.HandlerFunc(as.HomePage))
    r.HandleFunc(util.LOADNEWDATA, http.HandlerFunc(as.LoadNewCustomers))
    r.HandleFunc(util.GET_USERBYID, http.HandlerFunc(as.GetRecordByID))
    http.Handle(util.HOME, r)
	http.ListenAndServe(":8080", r)

}
