//db package
package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//NewDbInstance creates a new connection to a MongoDb instance using provided uri and credentials
func NewDbInstance(uri, username, password string) (*mongo.Client, error) {
	timeout := 10 * time.Second
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Connection to Mongo database is setted up")
	return client, nil
}
