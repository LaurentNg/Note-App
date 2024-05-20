package mongodb_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID        	 primitive.ObjectID  `bson:"_id,omitempty"`
    UserId    	 primitive.ObjectID  `bson:"userId" binding:"required"`
    CreatedDate  string  			 `bson:"createdDate" binding:"required"`
	ModifiedDate string  			 `bson:"modifiedDate" binding:"required"`
    Value  		 string  			 `bson:"value" binding:"required"`
}