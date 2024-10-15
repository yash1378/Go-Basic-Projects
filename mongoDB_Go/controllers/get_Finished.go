package controllers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFinished() ([]*Task, error) {
	filter := bson.D{
		primitive.E{Key: "completed", Value: true},
	}

	return filterTasks(filter)
}
