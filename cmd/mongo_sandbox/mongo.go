package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
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
}

func main() {
	b := Test{
		ID:   "1",
		Name: "Igor",
		Age:  19,
	}
	b1 := Test{
		ID:   "2",
		Name: "Andrey",
		Age:  24,
	}

	res, err := collection.InsertOne(ctx, b)

	if errHandler(err) != nil {
		fmt.Printf("error occurred: %v", err)
	}
	fmt.Println(res)

	res, err = collection.InsertOne(ctx, b1)
	if errHandler(err) != nil {
		fmt.Printf("error occurred: %v", err)
	}
	fmt.Println(res)

	filter := bson.M{"id": "2"}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var result []Test
	if err := cursor.All(ctx, &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}

type Test struct {
	ID   string `bson:"_id" json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func errHandler(err error) error {
	errStr := err.Error()
	switch {
	case strings.HasPrefix(errStr, "(DuplicateKey)"):
		return nil
	}
	return err
}
