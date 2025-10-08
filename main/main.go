package main

import (
	"OrderProject/entities"
	"OrderProject/handlers"
	"OrderProject/repository"
	"OrderProject/services"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

import _ "github.com/lib/pq"

var gormDB *gorm.DB

func InitDB() {
	var err error
	//так хардкодить нельзя. Лучше вывести в конфиги и читать через условный envViper
	connection := "user=postgres password=1234 dbname=goDB host=db port=5432 sslmode=disable"
	gormDB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	errTwo := gormDB.AutoMigrate(&entities.Customer{}, &entities.Order{}, &entities.CustomerSegments{})
	if errTwo != nil {
		log.Fatal("Failed to migrate database", errTwo)
	}

	fmt.Println("Successfully connected to database")
}

func CloseDB() {
	db, err := gormDB.DB()
	if err != nil {
		log.Fatal("Failed to close DB", err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal("Failed to close DB", err)
	}
}

// название метода лучше называть migrateUp что явно указывает что ты поднимаешь миграцию
func makeMigration() {
	//так хардкодить нельзя. Лучше вывести в конфиги и читать через условный envViper
	dsn := "postgres://postgres:1234@db:5432/goDB?sslmode=disable"
	m, err := migrate.New("file:///app/db/migrations", dsn)
	if err != nil {
		panic(err)
	}
	m.Up()
}

// название метода лучше называть migrateDown что явно указывает что ты поднимаешь миграцию
func backMigration() {
	//так хардкодить нельзя. Лучше вывести в конфиги и читать через условный envViper
	dsn := "postgres://postgres:1234@localhost:5432/gorillaFirstProjDB?sslmode=disable"
	m, err := migrate.New("file:///app/db/migrations", dsn)
	if err != nil {
		panic(err)
	}
	m.Down()
}

func main() {

	InitDB()
	makeMigration()
	defer CloseDB()

	var customerRepository repository.CustomerRepository
	customerRepo := repository.NewCustomerRepository(gormDB)
	customerRepository = customerRepo

	var orderRepository repository.OrderRepository
	orderRepo := repository.NewOrderRepository(gormDB)
	orderRepository = orderRepo

	var customerSegmentsRepository repository.CustomerSegmentsRepository
	customerSegmentsRepo := repository.NewCustomerSegmentsRepository(gormDB)
	customerSegmentsRepository = customerSegmentsRepo

	var customerServices services.CustomerService
	customerServ := services.NewCustomerService(customerRepository)
	customerServices = customerServ

	var orderServices services.OrderService
	orderServ := services.NewOrderServices(orderRepository)
	orderServices = orderServ

	var customerSegmentsServices services.CustomerSegmentsService
	customerSegmentsServ := services.NewCustomerSegmentsServices(customerSegmentsRepository)
	customerSegmentsServices = customerSegmentsServ

	var customerHandler handlers.CustomerHandler
	customerHand := handlers.NewCustomerHandler(customerServices)
	customerHandler = customerHand

	var orderHandler handlers.OrderHandler
	orderHand := handlers.NewOrderHandler(orderServices)
	orderHandler = orderHand

	var customerSegmentsHandler handlers.CustomerSegmentsHandler
	customerSegmentsHand := handlers.NewCustomerSegmentsHandler(customerSegmentsServices)
	customerSegmentsHandler = customerSegmentsHand

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.HandleCustomerGet).Methods(http.MethodGet)
	//для описания (POST,GET и тд) лучше использовать константы из пакета http как пример выше
	router.HandleFunc("/customers", customerHandler.HandleCustomerPost).Methods("POST")
	router.HandleFunc("/customers", customerHandler.HandleCustomerPut).Methods("PUT")
	router.HandleFunc("/customers", customerHandler.HandleCustomerDelete).Methods("DELETE")

	router.HandleFunc("/orders/{id}", orderHandler.HandleOrderGet).Methods("GET")
	router.HandleFunc("/orders", orderHandler.HandleOrderPost).Methods("POST")
	router.HandleFunc("/orders", orderHandler.HandleOrderPut).Methods("PUT")
	router.HandleFunc("/orders", orderHandler.HandleOrderDelete).Methods("DELETE")

	router.HandleFunc("/segments", customerSegmentsHandler.HandleCustomerSegmentsGet).Methods("GET")
	router.HandleFunc("/segments", customerSegmentsHandler.HandleCustomerSegmentsPost).Methods("POST")
	router.HandleFunc("/segments", customerSegmentsHandler.HandleCustomerSegmentsPut).Methods("PUT")
	router.HandleFunc("/segments", customerSegmentsHandler.HandleCustomerSegmentsDelete).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	errServer := server.ListenAndServe()
	if errServer != nil {
		log.Fatal("Error starting server:", errServer)
	}
}
