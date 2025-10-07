package handlers

import (
	"OrderProject/entities"
	"OrderProject/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type OrderHandler interface {
	HandleOrderGet(w http.ResponseWriter, r *http.Request)
	HandleOrderPost(w http.ResponseWriter, r *http.Request)
	HandleOrderPut(w http.ResponseWriter, r *http.Request)
	HandleOrderDelete(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	orderServices services.OrderService
}

func NewOrderHandler(orderServices services.OrderService) OrderHandler {
	return &orderHandler{orderServices: orderServices}
}

func (handler *orderHandler) HandleOrderGet(w http.ResponseWriter, r *http.Request) {
	//так делать нельзя, ибо ты теряешь идемпотентность сервиса
	//отправил айди, получил 1 запись, не отправил, получил массив
	//это плохая практика. Лучше явно создать 2 эндпойнта
	// .../orders [get]
	// .../order?order_id=123 [get]
	var idStr string = r.URL.Query().Get("id")
	if idStr == "" {
		err := json.NewEncoder(w).Encode(handler.orderServices.GetAllOrders())
		if err != nil {
			log.Fatal("Error getting all orders", err)
		}
	} else if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("Error converting order id to int", err)
		}
		errTwo := json.NewEncoder(w).Encode(handler.orderServices.GetByIdOrder(id))
		if errTwo != nil {
			log.Fatal("Error getting all orders", errTwo)
		}
	}
}

func (handler *orderHandler) HandleOrderPost(w http.ResponseWriter, r *http.Request) {
	var order entities.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatal("Error decoding body", err)
	}
	handler.orderServices.AddOrder(order)
}

func (handler *orderHandler) HandleOrderPut(w http.ResponseWriter, r *http.Request) {
	var order entities.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatal("Error decoding body", err)
	}
	handler.orderServices.UpdateOrder(order)
}

func (handler *orderHandler) HandleOrderDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal("Error converting order id to int", err)
	}
	handler.orderServices.DeleteOrder(id)
}
