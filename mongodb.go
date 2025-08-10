package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

var client *mongo.Client

// ConnectMongo connects to MongoDB
func ConnectMongo() {
	// context with timeout (if take longer time than this for )
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// connect to mongoDB
	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error in connecting to mongoDB: %v", err)
	}

	// to verify the connection use ping
	err = client.Ping(ctx, nil)	 
	if err != nil {
		log.Fatalf("Database doesn't Ping : %v", err)
	}

	fmt.Println("Database Connected successfully !")
}

func InsertCar(car Car) {
	collection := client.Database("database_name").Collection("Collection_name")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, car)
	if err != nil {
		log.Fatalf("Unable to insert car: %v", err)
	}

	fmt.Println("Car inserted successfully!")
	fmt.Println("Car insert successfully!")
}

func GetCarById(id int) {
	collection := client.Database("database_name").Collection("Collection_name")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	var car Car

	err := collection.FindOne(ctx, filter).Decode(&car)
	if err != nil {
		log.Fatalf("Error in getting car data : %v", err)
		return 
	}
	
	fmt.Printf("Car found: %+v", car)
}


func GenerateNextID() int {
	collection := client.Database("database_name").Collection("Collection_name")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	filter := bson.M{"_id": "carID"}
	update := bson.M{"$inc" : bson.M{"sequence":1}}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	
	var result struct {
		Sequence int `bson:"sequence"`
	}

	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		log.Fatalf("Error in generating new id: %v", err)
	}

	return result.Sequence
	return result.Sequence
}

func main() {
	ConnectMongo()
}