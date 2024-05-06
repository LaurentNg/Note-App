package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectMongoDb() (error) {

	url := ""
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	// defer func() {
    //     if err := client.Disconnect(context.TODO()); err != nil {
    //         panic(err)
    //     }
    // }()

	mongoClient = client
	fmt.Println("Connected to MongoDB!")
	return nil
}