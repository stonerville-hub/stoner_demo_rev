package handler

import (
	"net/http"
	"strings"
	"fmt"
)

func Controller() {
	getProductsHandler := http.HandlerFunc(getProducts)
	http.Handle("/user", getProductsHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Started successfully.")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// pathVar := r.URL.Path[len("/user/"):]
	query := r.URL.Query()
	filters, present := query["filters"] //filters=["color", "price", "brand"]
	if !present || len(filters) == 0 {
		fmt.Println("filters not present")
	}
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(filters, ",")))
}
