package handlers

import (
	"OrderProject/entities"
	"OrderProject/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CustomerSegmentsHandler interface {
	HandleCustomerSegmentsGet(w http.ResponseWriter, r *http.Request)
	HandleCustomerSegmentsPost(w http.ResponseWriter, r *http.Request)
	HandleCustomerSegmentsPut(w http.ResponseWriter, r *http.Request)
	HandleCustomerSegmentsDelete(w http.ResponseWriter, r *http.Request)
}

type customerSegmentsHandler struct {
	customerSegmentsServices services.CustomerSegmentsService
}

func NewCustomerSegmentsHandler(customerSegmentsServices services.CustomerSegmentsService) CustomerSegmentsHandler {
	return &customerSegmentsHandler{customerSegmentsServices: customerSegmentsServices}
}

func (handler *customerSegmentsHandler) HandleCustomerSegmentsGet(w http.ResponseWriter, r *http.Request) {
	var idStr string = r.URL.Query().Get("id")
	if idStr == "" {
		err := json.NewEncoder(w).Encode(handler.customerSegmentsServices.GetAllSegments())
		if err != nil {
			log.Fatal("Error getting customer segments from database", err)
		}
	} else if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("Error getting customer segments from database", err)
		}

		errTwo := json.NewEncoder(w).Encode(handler.customerSegmentsServices.GetByIdSegment(id))
		if errTwo != nil {
			log.Fatal("Error getting customer segments from database", errTwo)
		}
	}
}

func (handler *customerSegmentsHandler) HandleCustomerSegmentsPost(w http.ResponseWriter, r *http.Request) {
	var segments entities.CustomerSegments
	err := json.NewDecoder(r.Body).Decode(&segments)
	if err != nil {
		log.Fatal("Error getting customer segments from database", err)
	}
	handler.customerSegmentsServices.AddSegment(segments)
}

func (handler *customerSegmentsHandler) HandleCustomerSegmentsPut(w http.ResponseWriter, r *http.Request) {
	var segments entities.CustomerSegments
	err := json.NewDecoder(r.Body).Decode(&segments)
	if err != nil {
		log.Fatal("Error getting customer segments from database", err)
	}
	handler.customerSegmentsServices.UpdateSegment(segments)
}

func (handler *customerSegmentsHandler) HandleCustomerSegmentsDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal("Error getting customer segments from database", err)
	}
	handler.customerSegmentsServices.DeleteSegment(id)
}
