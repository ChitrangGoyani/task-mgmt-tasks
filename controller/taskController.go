package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/database"
	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTasks(c *fiber.Ctx) error {
	mg := &database.MG
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("tasks").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var tasks []models.Tasks = make([]models.Tasks, 0)
	// iterate through the results and decode each document
	if err := cursor.All(c.Context(), &tasks); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	mg := &database.MG
	collection := mg.Db.Collection("tasks")
	task := new(models.Tasks)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	task.ID = ""
	// insert task in mongodb collection
	insertResult, err := collection.InsertOne(c.Context(), task)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	// verify task insertion and return the inserted task
	query := bson.D{{Key: "_id", Value: insertResult.InsertedID}}
	cursor := collection.FindOne(c.Context(), query)
	createdTask := &models.Tasks{}
	cursor.Decode(createdTask)

	return c.Status(200).JSON(createdTask)
}

func UpdateTask(c *fiber.Ctx) error {
	mg := &database.MG
	collection := mg.Db.Collection("tasks")
	idParam := c.Params("id")
	taskId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(400)
	}
	var task models.Tasks
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: taskId}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "userID", Value: task.UserID},
			{Key: "priority", Value: task.Priority},
			{Key: "content", Value: task.Content},
			{Key: "time", Value: task.Time},
			{Key: "updatedTime", Value: task.UpdatedTime},
			{Key: "done", Value: task.Done},
		},
	}}
	err = collection.FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).SendString("Nahi mila yaar")
		}
		return c.SendStatus(500)
	}

	task.ID = idParam
	return c.Status(200).JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	mg := &database.MG
	collection := mg.Db.Collection("tasks")
	idParam := c.Params("id")
	taskId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: taskId}}
	deleteResult, err := collection.DeleteOne(c.Context(), query)
	if err != nil {
		c.SendStatus(500)
	}
	if deleteResult.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON("record deleted")
}

func SearchTask(c *fiber.Ctx) error {
	searchText := c.Params("content")
	mg := &database.MG
	collection := mg.Db.Collection("tasks")
	searchStage := bson.D{{Key: "$search", Value: bson.D{{Key: "index", Value: "search-task"}, {Key: "text", Value: bson.D{{Key: "path", Value: "content"}, {Key: "query", Value: searchText}, {Key: "fuzzy", Value: bson.D{{Key: "maxEdits", Value: 2}}}}}}}}
	// the amount of time the operation can run on the server
	opts := options.Aggregate().SetMaxTime(60 * time.Second)
	cursor, err := collection.Aggregate(c.Context(), mongo.Pipeline{searchStage}, opts)
	if err != nil {
		return c.SendString(err.Error())
	}
	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		return c.SendStatus(500)
	}
	fmt.Println(results)
	return c.Status(200).JSON(results)
}
