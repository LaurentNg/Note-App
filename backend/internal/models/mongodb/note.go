package mongodb_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        	 primitive.ObjectID  `bson:"_id,omitempty"`
    UserId    	 primitive.ObjectID  `bson:"userId" binding:"required"`
    CreatedDate  time.Time  		 `bson:"createdDate" binding:"required"`
	ModifiedDate time.Time  	     `bson:"modifiedDate" binding:"required"`
    Value  		 string  			 `bson:"value" binding:"required"`
}