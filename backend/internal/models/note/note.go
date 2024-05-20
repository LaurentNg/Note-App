package note

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID        	 primitive.ObjectID  `json:"_id,omitempty"`
    UserId    	 primitive.ObjectID  `json:"userId" binding:"required"`
    CreatedDate  string  			 `json:"createdDate" binding:"required"`
	ModifiedDate string  			 `json:"modifiedDate" binding:"required"`
    Value  		 string  			 `json:"value" binding:"required"`
}