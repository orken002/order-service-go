package handlers

import (
	"OrderProject/entities"
	"OrderProject/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CustomerHandler interface {
	HandleCustomerGet(w http.ResponseWriter, r *http.Request)
	HandleCustomerPost(w http.ResponseWriter, r *http.Request)
	HandleCustomerPut(w http.ResponseWriter, r *http.Request)
	HandleCustomerDelete(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) CustomerHandler {
	return &customerHandler{customerService: customerService}
}

func (handler customerHandler) HandleCustomerGet(w http.ResponseWriter, r *http.Request) {
	var idStr string = r.URL.Query().Get("id")
	var nameStr string = r.URL.Query().Get("name")

	if idStr == "" || nameStr == "" {
		err := json.NewEncoder(w).Encode(handler.customerService.GetAllCustomers())
		if err != nil {
			log.Fatal("Error getting customers from database", err)
		} else if idStr != "" && nameStr == "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Fatal("Error getting customers from database", err)
			}
			errTwo := json.NewEncoder(w).Encode(handler.customerService.GetByIdCustomer(id))
			if errTwo != nil {
				log.Fatal("Error getting customers from database", errTwo)
			}
		} else if idStr == "" && nameStr != "" {
			err := json.NewEncoder(w).Encode(handler.customerService.GetByNameCustomer(nameStr))
			if err != nil {
				log.Fatal("Error getting customers from database", err)
			}
		}
	}
}

func (handler customerHandler) HandleCustomerPost(w http.ResponseWriter, r *http.Request) {
	var customer entities.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Fatal("Error getting customers from database", err)
	}
	handler.customerService.AddCustomer(customer)
}

func (handler customerHandler) HandleCustomerPut(w http.ResponseWriter, r *http.Request) {
	var customer entities.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Fatal("Error getting customers from database", err)
	}
	handler.customerService.UpdateCustomer(customer)
}

func (handler customerHandler) HandleCustomerDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal("Error getting customers from database", err)
	}
	handler.customerService.DeleteCustomer(id)
}
