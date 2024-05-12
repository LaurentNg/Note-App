package mongodb

import (
	"Note-App/internal/models"
	"Note-App/internal/services/logger"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrExistingUser = errors.New("user already exists")
)

func CreateUser(newUser models.User) error {
	logger.Info(fmt.Sprintf("Creating user with email: %s and username: %s", newUser.Email, newUser.Username))
	
	// Check if user already exists
	if err := checkExistingUser(newUser); err == nil {
		logger.Error(fmt.Sprintf("Duplicate : user with email: %s and username: %s already exist", newUser.Email, newUser.Username))
		return ErrExistingUser
	}

	coll := mongoClient.Database("notedb").Collection("users")
	userBSON, err := bson.Marshal(newUser)
	if err != nil {
		return err
	}

	_, err = coll.InsertOne(context.TODO(), userBSON)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating user with email: %s and username: %s", newUser.Email, newUser.Username))
		return err
	}

	logger.Info(fmt.Sprintf("User with email: %s and username: %s created successfully", newUser.Email, newUser.Username))
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
