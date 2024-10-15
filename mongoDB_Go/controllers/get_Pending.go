package controllers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPending() ([]*Task, error) {
	filter := bson.D{
		primitive.E{Key: "completed", Value: false},
	}

	return filterTasks(filter)
}
