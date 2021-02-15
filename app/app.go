package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//define routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customer/{customer_id}", getCustomer)

	//starting server
	err := http.ListenAndServe("localhost:4000", router)

	if err != nil {
		log.Fatal(err)
	}
}
