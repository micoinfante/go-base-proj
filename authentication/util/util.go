package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func toObjectId(id string) primitive.ObjectID {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	return objectId
}