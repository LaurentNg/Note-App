package mongodb

import (
	"Note-App/internal/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrExistingUser = errors.New("user already exists")
)

func CreateUser(newUser models.User) error {
	fmt.Println("Creating user")

	// Check if user already exists
	if err := checkExistingUser(newUser); err == nil {
		return ErrExistingUser
	}

	coll := mongoClient.Database("notedb").Collection("users")
	userBSON, err := bson.Marshal(newUser)
	if err != nil {
		return err
	}

	_, err = coll.InsertOne(context.TODO(), userBSON)
	if err != nil {
		return err
	}

	fmt.Println("Successfully created user")
	return nil
}

func checkExistingUser(userToCheck models.User) error {
	coll := mongoClient.Database("notedb").Collection("users")

	var user models.User
	filter := bson.M{
		"$or": bson.A{
			bson.M{"email": userToCheck.Email},
			bson.M{"username": userToCheck.Username},
		},
	}

	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	return err
}
