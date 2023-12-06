package collection

import "go.mongodb.org/mongo-driver/mongo"

func GetCollection(client *mongo.Client, db string, collectionName string) *mongo.Collection {
	collection := client.Database(db).Collection(collectionName)
	return collection
}
