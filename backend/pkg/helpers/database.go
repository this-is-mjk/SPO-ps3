package helpers

import (
	"context"
	"log"
	"os"

	"github.com/this-is-mjk/SPO/p3/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Collection {
	MONGO_URI := os.Getenv("MONGO_URI")
	// Connect to the database.
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Create collection
	collection := client.Database("SPO-P3").Collection("users")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return collection

}
func FindUser(users *mongo.Collection, filter interface{}) (models.User, bool) {
	result := users.FindOne(context.Background(), filter)
	var user models.User
	err := result.Decode(&user)
	if err != nil {
		return models.User{}, false
	}
	return user, true
}
func InsertUser(users *mongo.Collection, user interface{}) (*mongo.InsertOneResult, error) {
	return users.InsertOne(context.Background(), user)
}
