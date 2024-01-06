package repository

import (
	"context"
	"log"

	"github.com/lucas-code42/testMongoDb/pkg/api/config"
	"github.com/lucas-code42/testMongoDb/pkg/api/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var repositoryContext = context.TODO()

type BookRepositoryImp interface {
	GetBook() []entity.Book
	CreateBook(book entity.Book) error
	UpdateBook(book entity.Book) error
	DeleteBook(book entity.Book) error
}

type MongoBookRepository struct {
	DataBaseConnection *mongo.Database
}

func NewBookRepository(dbType string) BookRepositoryImp {
	switch dbType {
	case "mongo":
		client := mongoDb{}
		return &MongoBookRepository{
			DataBaseConnection: client.connectToMongoDb(),
		}
	case "redis":
		// connect to redis...
		panic("not implemented")
	default:
		panic("NewBookRepository just apply redis or mongo")
	}
}

type mongoDb struct{}

func (m *mongoDb) connectToMongoDb() *mongo.Database {
	mongoSettings := config.GetMongoConfig()

	opts := options.Client().ApplyURI(
		mongoSettings.Uri,
	)

	client, err := mongo.Connect(repositoryContext, opts)
	if err != nil {
		log.Panicf("[*] error to connect with mongoDb.\n%v", err.Error())
	}

	if m.pingMongoDb(client) != nil {
		log.Panicf("[*] error to ping mongoDb.\n%v", err.Error())
	}

	mongoDabatase := client.Database(mongoSettings.DefaultDb)
	return mongoDabatase
}

func (m *mongoDb) pingMongoDb(client *mongo.Client) error {
	return client.Ping(repositoryContext, nil)
}

// apenas para ilustrar exemplo
type redisDb struct{}

func (r *redisDb) connectToRedisDb() error {
	panic("not implemented")
}

func (r *redisDb) pingRedisDb() error {
	panic("not implemented")
}
