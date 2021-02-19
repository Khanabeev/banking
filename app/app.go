package app

import (
	"fmt"
	"github.com/Khanabeev/banking/domain"
	"github.com/Khanabeev/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start() {

	router := mux.NewRouter()

	//wiring all together
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)

	if err != nil {
		log.Fatal(err)
	}
}
