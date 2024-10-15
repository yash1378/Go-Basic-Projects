package controllers

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTask(text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}

	res, err := Collection.DeleteOne(Ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No tasks were deleted")
	}

	return nil
}
