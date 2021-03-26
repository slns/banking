package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slns/banking/app/domain"
	"github.com/slns/banking/app/service"
)


func Start()  {

	// mux := http.NewServeMux()
	 router := mux.NewRouter()

	 //wiring
	//  ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	 ch := CustomerHandlers{service.NewCustomerService(domain.NewcustomerrepositoryDb())}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//Starting Server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}