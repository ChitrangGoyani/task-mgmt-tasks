package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var dbName = "trial"
var MG = MongoInstance{}

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoPass := os.Getenv("MONGO_USER_PASS")
	uri := fmt.Sprintf("mongodb+srv://cgoyani:%s@trial.cek3scp.mongodb.net/?retryWrites=true&w=majority", mongoPass)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic("Could not connect to mongodb database")
	}
	// defer context.WithCancel(context.TODO())
	db := client.Database(dbName)
	MG.Client = client
	MG.Db = db

	return nil
}
