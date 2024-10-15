package controllers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CompleteTask(text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	//as t is not getting used anywhere so can be removed easily for better practice
	// t := &Task{}
	// return Collection.FindOneAndUpdate(Ctx, filter, update).Decode(t)
	result := Collection.FindOneAndUpdate(Ctx, filter, update)

	// Check if an error occurred
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
