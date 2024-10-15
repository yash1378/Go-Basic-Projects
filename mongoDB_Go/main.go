package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"tasker/controllers"

	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// added the mongo and options packages, which the MongoDB Go driver provides.

func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://yashdugriyal1066:mongo123456@cluster0.phiye.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	client, err := mongo.Connect(controllers.Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(controllers.Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	controllers.Collection = client.Database("tasker").Collection("tasks")

	// Your Go program connects to a MongoDB Atlas cluster, but it does not create the database and controllers.Collection explicitly. MongoDB creates a database and controllers.Collection automatically when you first write data to them.

	// // Insert a sample document to create the database and controllers.Collection
	// sampleTask := bson.M{"name": "Sample Task", "completed": false}
	// _, err = controllers.Collection.InsertOne(ctx, sampleTask)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Sample task inserted into the controllers.Collection.")
}

func main() {
	app := &cli.App{
		Name:  "tasker",
		Usage: "A simple CLI program to manage your tasks",
		Action: func(c *cli.Context) error {
			tasks, err := controllers.GetPending()
			if err != nil {
				if err == mongo.ErrNoDocuments {
					fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
					return nil
				}

				return err
			}

			controllers.PrintTasks(tasks)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					if str == "" {
						return errors.New("Cannot add an empty task")
					}

					task := &controllers.Task{
						ID:        primitive.NewObjectID(),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Text:      str,
						Completed: false,
					}

					return controllers.CreateTask(task)
				},
			},
			{
				Name:    "all",
				Aliases: []string{"l"},
				Usage:   "list all tasks",
				Action: func(c *cli.Context) error {
					tasks, err := controllers.GetAll()
					if err != nil {
						if err == mongo.ErrNoDocuments {
							fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
							return nil
						}

						return err
					}
					controllers.PrintTasks(tasks)
					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					text := c.Args().First()
					return controllers.CompleteTask(text)
				},
			},
			{
				Name:    "finished",
				Aliases: []string{"f"},
				Usage:   "list completed tasks",
				Action: func(c *cli.Context) error {
					tasks, err := controllers.GetFinished()
					if err != nil {
						if err == mongo.ErrNoDocuments {
							fmt.Print("Nothing to see here.\nRun `done 'task'` to complete a task")
							return nil
						}

						return err
					}

					controllers.PrintTasks(tasks)
					return nil
				},
			},
			{
				Name:  "rm",
				Usage: "deletes a task on the list",
				Action: func(c *cli.Context) error {
					text := c.Args().First()
					err := controllers.DeleteTask(text)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
