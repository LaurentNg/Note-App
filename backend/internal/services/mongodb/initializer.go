package mongodb

import (
	"Note-App/internal/services/logger"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectMongoDb() (error) {

	url := ""
	clientOptions := options.Client().ApplyURI(url)

	logger.Info("Connecting to MongoDB...")
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
	logger.Info("Connected to MongoDB")
	return nil
}