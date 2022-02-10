package constants

import (
	"strings"

	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const (
	uriMongo        = "mongodb://user:pass@localhost:2717/?maxPoolSize=20&w=majority"
	mongoDatabase   = "myStuff"
	mongoCollection = "test"
)

func GenerateApiKey() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func RetrieveMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uriMongo))
	if err != nil {
		panic(err)
	}
	return client
}

func RetrieveMongoCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(mongoDatabase).Collection(mongoDatabase)
}
