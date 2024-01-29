package app

import (
	"banking/domain"
	"banking/handler"
	"banking/logger"
	"banking/service"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func Start() {
	router := mux.NewRouter()

	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	ch := handler.NewCustomerHandler(service.NewCustomerService(customerRepositoryDb))

	//customerRepositoryStub := domain.NewCustomerRepositoryStub()
	//ch := handler.NewCustomerHandler(service.NewCustomerService(customerRepositoryStub))

	router.
		HandleFunc("/customers", ch.GetAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id}", ch.GetCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers", ch.CreateCustomer).
		Methods(http.MethodPost).
		Name("CreateCustomer")

	logger.Info("Starting the Application...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDbClient() *mongo.Client {
	ctx := context.Background()
	mongoDbUri := "mongodb://localhost:27020/"
	clientOptions := options.Client().ApplyURI(mongoDbUri).
		SetMaxPoolSize(20).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(time.Minute * 3).
		SetDirect(true)
	if logger.IsDebugEnabled() {
		logger.Debug("getDbClient() | Attempting to connect to Database ")
	}
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	if logger.IsDebugEnabled() {
		logger.Debug("getDbClient() |  Connected to Database ")
	}
	return client
}
