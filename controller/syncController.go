package controller

import (
	"context"
	"fmt"

	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/database"
	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/kafka"
	"go.mongodb.org/mongo-driver/mongo"
)

// stream changes of mongodb data to kafka and then from kafka send it over to elastic search
var OpenStream *mongo.ChangeStream

func OpenChangeStream() {
	collection := database.MG.Db.Collection("tasks")
	changeStream, err := collection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	// Iterates over the cursor to print the change stream events
	for changeStream.Next(context.TODO()) {
		fmt.Println(changeStream.Current)
		kafka.Produce("tasks", 0, []byte("Test")) // send event update to kafka
	}
	OpenStream = changeStream
}

func CloseChangeStream() {
	OpenStream.Close(context.TODO())
}
