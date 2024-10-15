package controllers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gookit/color.v1"
)

func GetAll() ([]*Task, error) {
	filter := bson.D{{}}
	return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*Task, error) {
	var tasks []*Task

	cur, err := Collection.Find(Ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(Ctx) {
		var t Task
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	// once exhausted, close the cursor
	cur.Close(Ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func PrintTasks(tasks []*Task) {
	for i, v := range tasks {
		if v.Completed {
			color.Green.Printf("%d: %s\n", i+1, v.Text)
		} else {
			color.Red.Printf("%d: %s\n", i+1, v.Text)
		}
	}
}
