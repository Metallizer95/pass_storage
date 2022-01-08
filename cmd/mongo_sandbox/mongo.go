package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasker")
	type f struct {
		D string `json:"D"`
	}
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	r, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.InsertedID)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func main() {

}
