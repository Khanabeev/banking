package app

import (
	"github.com/Khanabeev/banking/domain"
	"github.com/Khanabeev/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//wiring all together
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//starting server
	err := http.ListenAndServe("localhost:4000", router)

	if err != nil {
		log.Fatal(err)
	}
}
