package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var BalancesCollection *mongo.Collection

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	MongoClient = client
	BalancesCollection = client.Database(os.Getenv("MONGO_DB")).Collection("balances")
}

func SaveBalance(ctx context.Context, address, balance string) {
	filter := bson.M{"address": address}
	update := bson.M{"$set": bson.M{"balance": balance}}
	_, err := BalancesCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		log.Println("Mongo update error:", err)
	}
}
