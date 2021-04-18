package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"tasky/model"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	globalClient *mongo.Client
)

type MongoRepository struct{}

type NewMongoRepository interface {
	GetCollection(collection string) (*mongo.Collection, *model.Error)
	Disconnect()
}

func New() *MongoRepository {
	return &MongoRepository{}
}

func (m *MongoRepository) Disconnect() {
	if globalClient == nil {
		return
	}

	err := globalClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database successfully disconnected")
}

func (m *MongoRepository) GetCollection(collection string) (*mongo.Collection, *model.Error) {
	usr := viper.Get("DEV_MONGODB_USER")
	pwd := viper.Get("DEV_MONGODB_PASSWORD")
	database := viper.Get("DEV_MONGODB_DATABASE")
	uri := fmt.Sprintf("mongodb+srv://%s:%s@testgo.6itcv.mongodb.net/%s?retryWrites=true&w=majority",
		usr.(string), pwd.(string), database.(string))

	globalClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		tempErr := &model.Error{
			Message: err.Error(),
		}
		fmt.Println(&tempErr.Message)
		return nil, tempErr
	}

	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = globalClient.Connect(context)
	if err != nil {
		tempErr := &model.Error{
			Message: err.Error(),
		}
		fmt.Println(&tempErr.Message)
		return nil, tempErr
	}

	databases, err := globalClient.ListDatabaseNames(context, bson.M{})
	if err != nil {
		tempErr := &model.Error{
			Message: err.Error(),
		}
		fmt.Println(&tempErr.Message)
		return nil, tempErr
	}
	fmt.Println("Databases: ", databases)

	return globalClient.Database(database.(string)).Collection(collection), nil
}
