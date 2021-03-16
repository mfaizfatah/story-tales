package adapter

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoClient *mongo.Client
)

// LoadMySQL is load connection to mysql server
func LoadMongoDB(url string) {
	mongoClient = MongoDB(url)
}

// MySQL is open connection to mysql server
func MongoDB(url string) *mongo.Client {
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client
}

// DBSQL is open connection into database
func Mongo() *mongo.Client {
	return mongoClient
}

//DB is func
func MongoDatabase() *mongo.Database {
	return mongoClient.Database(os.Getenv("MONGO_DB"))
}
