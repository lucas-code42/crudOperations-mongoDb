package api

import (
	"context"
	"log"

	"github.com/lucas-code42/testMongoDb/pkg/api/controller"
	"github.com/lucas-code42/testMongoDb/pkg/api/repository"
	"github.com/lucas-code42/testMongoDb/pkg/api/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri               = "mongodb://root:example@mongo:27017/"
	defaultDb         = "db"
	defaultCollection = "books"
)

var ctx = context.TODO()

func InitDependencies() controller.ControllerImp {
	mongoDataBaseClient := connectToMongoDb()
	
	// TODO: criar uma camada de banco mais abstrata??
	// camada de baixo nivel
	repo := repository.NewBookRepository( mongoDataBaseClient)

	// camada de dominio
	service := service.NewBookService(repo)

	// camada da api/controller
	controller := controller.NewController(service)
	return controller
}

func connectToMongoDb() *mongo.Client {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Panicf("[*] error to connect with mongoDb.\n%v", err.Error())
	}

	if pingMongoDb(client) != nil {
		log.Panicf("[*] error to ping mongoDb.\n%v", err.Error())
	}
	return client
}

func pingMongoDb(client *mongo.Client) error {
	return client.Ping(ctx, nil)
}
