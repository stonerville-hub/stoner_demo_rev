package mongoCRUD

import (
	"context"
	"fmt"

	"github.com/my/repo/server/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InsertRecord() {
	fmt.Println("----Start - Insert Record----")
	// Create a new client, connect to the server, and return collection
	client := constants.RetrieveMongoClient()
	coll := constants.RetrieveMongoCollection(client)

	doc := bson.D{{"api_key", constants.GenerateApiKey()}, {"first_name", "Italo Calvino"}, {"last_name", 1974}, {"company", 1974}}
	result, err := coll.InsertOne(context.TODO(), doc)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Println("----END - Insert Record----")
}

func CheckDBConnection() {
	fmt.Println("----Start - Check DB Connection----")
	// Create a new client and connect to the server
	client := constants.RetrieveMongoClient()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Println("----END - Check DB Connection----")
}
