// repo.go
package repository

import (
	"errors"
	"time"

	"github.com/JeerasakTH/go-test-crud/collection"
	"github.com/JeerasakTH/go-test-crud/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func CreateOne(dbName string, collectionName string, payload interface{}) (*mongo.InsertOneResult, error) {
	DB, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	getCollection := collection.GetCollection(DB, dbName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.InsertOne(ctx, payload)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func CreateMany(dbName string, collectionName string, payload []interface{}) (*mongo.InsertManyResult, error) {
	DB, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	getCollection := collection.GetCollection(DB, dbName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.InsertMany(ctx, payload)
	if err != nil {
		return obj, err
	}

	if obj == nil {
		err := errors.New("INSERT FAIL")
		return obj, err
	}
	return obj, nil
}

func GetOne(dbName string, collectionName string, filter interface{}, filterOption interface{}, data interface{}) error {
	DB, err := database.ConnectDB()
	if err != nil {
		return err
	}

	getCollection := collection.GetCollection(DB, dbName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	option := options.FindOne()
	option.SetSort(filterOption)
	defer cancel()

	err = getCollection.FindOne(ctx, filter, option).Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func GetMany(dbName string, collectionName string, filter interface{}, filterOption interface{}, data interface{}) error {
	DB, err := database.ConnectDB()
	if err != nil {
		return err
	}

	getCollection := collection.GetCollection(DB, dbName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	option := options.Find()
	option.SetSort(filterOption)
	defer cancel()

	obj, err := getCollection.Find(ctx, filter, option)
	if err != nil {
		return err
	}

	err = obj.All(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
