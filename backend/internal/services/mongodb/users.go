package mongodb

import (
	mongodb_models "Note-App/internal/models/mongodb"
	"Note-App/internal/services/logger"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrExistingUser = errors.New("user already exists")
)

func CreateUser(newUser *mongodb_models.User) error {
	logger.Info(fmt.Sprintf("Creating user with email: %s and username: %s", newUser.Email, newUser.Username))
	
	// Check if user already exists
	if err := checkExistingUser(newUser); err == nil {
		logger.Error(fmt.Sprintf("Duplicate : user with email: %s and username: %s already exist", newUser.Email, newUser.Username))
		return ErrExistingUser
	}

	// Hash password
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		logger.Error(fmt.Sprintf("Error hashing password for user with email: %s and username: %s", newUser.Email, newUser.Username))
		return err
	}

	newUser.Password = hashedPassword

	// Insert user into database
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

func GetUserByEmail(email string) (mongodb_models.User, error) {
	coll := mongoClient.Database("notedb").Collection("users")

	var user mongodb_models.User
	filter := bson.M{"email": email}

	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return mongodb_models.User{}, err
	}

	return user, nil
}

func checkExistingUser(userToCheck *mongodb_models.User) error {
	coll := mongoClient.Database("notedb").Collection("users")

	var user mongodb_models.User
	filter := bson.M{
		"$or": bson.A{
			bson.M{"email": userToCheck.Email},
			bson.M{"username": userToCheck.Username},
		},
	}

	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	return err
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
