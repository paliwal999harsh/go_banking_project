package domain

import (
	"banking/errs"
	"banking/logger"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepositoryDb struct {
	client *mongo.Client
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	var err error
	var customers []Customer
	coll := d.client.Database("banking").Collection("customer")
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Debug("FindAll() | No Documents Found")
			return nil, errs.NewNotFoundError("No Records Found")
		}
	}
	if err = cursor.All(context.TODO(), &customers); err != nil {
		logger.Error("FindAll() | Error While cursoring the data " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error Occurred")
	}
	// Prints the results of the find operation as structs
	for _, result := range customers {
		_ = cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		logger.Debug(string(output))
	}
	return customers, nil
}
func (d CustomerRepositoryDb) FindAllByStatus(status bool) ([]Customer, *errs.AppError) {
	var err error
	var customers []Customer

	coll := d.client.Database("banking").Collection("customer")
	filter := bson.D{{"status", status}}
	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Debug("FindAllByStatus() | No Documents Found")
			return nil, errs.NewNotFoundError("No Records Found")
		}
	}
	if err = cursor.All(context.TODO(), &customers); err != nil {
		logger.Error("Error While cursoring the data " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error Occurred")
	}

	// Prints the results of the find operation as structs
	for _, result := range customers {
		_ = cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		logger.Debug(string(output))
	}
	return customers, nil
}
func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {

	var err error
	var customer Customer

	coll := d.client.Database("banking").Collection("customer")
	filter := bson.D{{"_id", id}}
	err = coll.FindOne(context.TODO(), filter).Decode(&customer)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Debug("FindById() | No Documents Found")
			return nil, errs.NewNotFoundError("No Records Found")
		}
	}
	return &customer, nil
}

func (d CustomerRepositoryDb) Save(c Customer) (**mongo.InsertOneResult, *errs.AppError) {
	coll := d.client.Database("banking").Collection("customer")
	result, err := coll.InsertOne(context.TODO(), c)
	if err != nil {
		logger.Error("Save() | Unable to save the document")
		if mongo.IsDuplicateKeyError(err) {
			logger.Error("Save() | Record already exsits with the given Customer ID")
			return nil, errs.NewBadRequestError("Record already exsits with the given Customer ID.")
		}
		return nil, errs.NewBadRequestError("Specified Format is not correct	.")
	}
	return &result, nil
}

func NewCustomerRepositoryDb(dbClient *mongo.Client) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
