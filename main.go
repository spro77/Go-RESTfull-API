package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Product prototype for our store
type Product struct {
	ID    string `json:id`
	Title string `json:title`
	Price string `json:price`
	Descr string `json:descr`
}

var store []Product

func main() {
	router := mux.NewRouter()

	store = append(store,
		Product{ID: "1", Title: "Elephant T-Shirt", Price: "9.99", Descr: "T-Shirt with the beautifull elephant print"},
		Product{ID: "2", Title: "Baseball Cap", Price: "4.99", Descr: "Cap with beer cans handles"},
		Product{ID: "3", Title: "White Office Shirt", Price: "19.99", Descr: "You are Boss"},
		Product{ID: "4", Title: "Blue Office Shirt", Price: "14.99", Descr: "You are like a Boss"},
		Product{ID: "5", Title: "untittled", Price: "0", Descr: ""},
	)

	router.HandleFunc("/store", getProducts).Methods("GET")
	router.HandleFunc("/store/{id}", getProduct).Methods("GET")
	router.HandleFunc("/store", addProduct).Methods("POST")
	router.HandleFunc("/store", updateProduct).Methods("PUT")
	router.HandleFunc("/store/{id}", removeProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(store)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	for _, Product := range store {
		if Product.ID == params["id"] {
			json.NewEncoder(w).Encode(&Product)
		}
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var Product Product
	_ = json.NewDecoder(r.Body).Decode(&Product)

	store = append(store, Product)
	json.NewEncoder(w).Encode(store)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	var Product Product
	_ = json.NewDecoder(r.Body).Decode(&Product)

	for i, item := range store {
		if item.ID == Product.ID {
			store[i] = Product
		}
	}

	json.NewEncoder(w).Encode(store)
}

func removeProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var Product Product

	Product.ID = params["id"]

	for i, item := range store {
		if item.ID == Product.ID {
			store = append(store[:i], store[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(store)
}
