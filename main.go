package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title  string
	Author string
}

const (
	uri               = "mongodb://root:example@mongo:27017/"
	defaultDb         = "db"
	defaultCollection = "books"
)

var ctx = context.TODO()

func connect() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Panicf("[*] error to connect with mongoDb.\n%v", err.Error())
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Panicf("[*] error to ping mongoDb.\n%v", err.Error())
	}
	return client, nil
}

func insertRecord(client *mongo.Client, data Book) error {
	coll := client.Database(defaultDb).Collection(defaultCollection)
	result, err := coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	fmt.Printf("[*] Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func insertMany(client *mongo.Client, data []interface{}) error {
	coll := client.Database(defaultDb).Collection(defaultCollection)
	result, err := coll.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}
	fmt.Printf("[*] Documents inserted: %v\n", len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		fmt.Printf("[*] Inserted document with _id: %v\n", id)
	}
	return nil
}

func queryFilter(client *mongo.Client) error {
	//                 key,               value
	filter := bson.D{{"title", "In Memory of Memory"}}
	coll := client.Database(defaultDb).Collection(defaultCollection)
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	var results []Book
	if err = cursor.All(ctx, &results); err != nil {
		return err
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}
	return nil
}

func updateData(client *mongo.Client) error {
	coll := client.Database(defaultDb).Collection(defaultCollection)
	filter := bson.D{{"title", "In Memory of Memory"}}
	update := bson.D{{"$set", bson.D{{"Title", "Harry Potter"}}}}
	result, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
	return nil
}

func deleteByFilter(client *mongo.Client) error {
	coll := client.Database(defaultDb).Collection(defaultCollection)
	filter := bson.D{{"title", "In Memory of Memory"}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)
	return nil
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")

	// based on https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/
	mongoClient, err := connect()
	if err != nil {
		log.Panicf(
			"[*] Something got wrong check the 'connect' func.\n%v", err.Error())
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Panicf(
				"[*] Troubles to disconnect dataBase.\n%v", err.Error())
		}
	}()

	// create
	if err = insertRecord(mongoClient, Book{Title: "Atonement", Author: "Ian McEwan"}); err != nil {
		log.Panicf(
			"[*] Something got wrong check the 'insertRecord' func.\n%v", err.Error())
	}
	// data example
	docs := []interface{}{
		Book{Title: "Cat's Cradle", Author: "Kurt Vonnegut Jr."},
		Book{Title: "In Memory of Memory", Author: "Maria Stepanova"},
		Book{Title: "Pride and Prejudice", Author: "Jane Austen"},
	}
	if err = insertMany(mongoClient, docs); err != nil {
		log.Panicf(
			"[*] Something got wrong check the 'insertMany' func.\n%v", err.Error())
	}

	// read
	if err = queryFilter(mongoClient); err != nil {
		log.Panicf(
			"[*] Something got wrong check the 'queryFilter' func.\n%v", err.Error())
	}

	// update
	if err = updateData(mongoClient); err != nil {
		log.Panicf(
			"[*] Something got wrong check the 'updateData' func.\n%v", err.Error())
	}

	// delete
	if err = deleteByFilter(mongoClient); err != nil {
		log.Panicf(
			"[*] Something got wrong check the 'deleteByFilter' func.\n%v", err.Error())
	}

	w.Write([]byte("Done"))

}

func main() {
	fmt.Println("init")
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
