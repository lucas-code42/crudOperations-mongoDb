package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title  string
	Author string
}

const (
	uri               = "mongodb://root:example@localhost:27017/"
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

func main() {
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
}
