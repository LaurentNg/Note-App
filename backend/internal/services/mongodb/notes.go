package mongodb

import (
	mongodb_errors "Note-App/internal/errors/mongodb"
	mongodb_models "Note-App/internal/models/mongodb"
	"Note-App/internal/services/logger"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(newNote *mongodb_models.Note) error {
	coll := mongoClient.Database("notedb").Collection("notes")

	noteBSON, err := bson.Marshal(newNote)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = coll.InsertOne(context.TODO(), noteBSON)
	if err != nil {
		err := mongodb_errors.ErrCreateNote()
		logger.Error(err.Error())
		return err
	}

	return nil
}

func GetNotesByUserId(userId primitive.ObjectID) ([]mongodb_models.Note, error) {
	coll := mongoClient.Database("notedb").Collection("notes")

	var notes []mongodb_models.Note
	filter := bson.M{"userId": userId}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var note mongodb_models.Note
		err := cursor.Decode(&note)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func DeleteNoteById(noteId primitive.ObjectID) error {
	coll := mongoClient.Database("notedb").Collection("notes")

	filter := bson.M{"_id": noteId}

	_,	err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		logger.Error(err.Error())
		return err
	}	

	return nil
}