package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"strings"
)

var mainCollection *mongo.Collection
var subCollection *mongo.Collection
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

	mainCollection = client.Database("tasker").Collection("main")
	subCollection = client.Database("tasker").Collection("sub")
}

func main() {
	addresses := createAddresses()
	for _, address := range addresses {
		_, err := subCollection.InsertOne(context.TODO(), address)
		if errHandler(err) != nil {
			fmt.Println(err)
		}
	}
	person := mainStruct{
		ID:          "1",
		Name:        "Ivan",
		Gender:      "Male",
		City:        "Saint-Petersburg",
		Country:     "Russia",
		AddressesID: []string{"1", "3", "4"},
	}
	_, err := mainCollection.InsertOne(ctx, person)
	if errHandler(err) != nil {
		fmt.Println(err)
	}

	filter := bson.M{"_id": "1"}
	cursor, err := mainCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var result []mainStruct
	if err := cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	if len(result) == 0 {
		panic("There is not result")
	}
	filter = bson.M{"_id": bson.M{"$in": result[0].AddressesID}}
	fmt.Println(filter)
	subCursor, err := subCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var a []Address
	if err := subCursor.All(ctx, &a); err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func createAddresses() []Address {
	var result []Address
	for i := 0; i < 10; i++ {
		id := strconv.Itoa(i)
		result = append(result, Address{
			ID:       id,
			Street:   "Kolomyazhskiy",
			Building: id,
		})
	}
	return result
}

type mainStruct struct {
	ID          string   `bson:"_id" json:"id"`
	Name        string   `json:"name"`
	Gender      string   `json:"gender"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	AddressesID []string `json:"addresses_id"`
}

type Address struct {
	ID       string `bson:"_id" json:"id"`
	Street   string `json:"street"`
	Building string `json:"building"`
}

func errHandler(err error) error {
	if err == nil {
		return nil
	}
	errStr := err.Error()
	switch {
	case strings.Contains(errStr, "duplicate"):
		return nil
	}
	return err
}
