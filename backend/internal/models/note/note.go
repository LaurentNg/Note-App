package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        	 primitive.ObjectID  `json:"_id,omitempty"`
    UserId    	 primitive.ObjectID  `json:"userId,omitempty"`
	CreatedDate  time.Time 			 `json:"createdDate,omitempty"`
	ModifiedDate time.Time  	     `json:"modifiedDate,omitempty"`
    Value  		 string  			 `json:"value" binding:"required"`
}