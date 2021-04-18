package connection

import (
	"context"
	"fmt"
	"time"

	"tasky/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct{}

type NewMongoRepository interface {
	GetCollection(collection string) (*mongo.Collection, *model.Error)
}

func New() *MongoRepository {
	return &MongoRepository{}
}

var (
	usr      = "root"
	pwd      = "Inception2575.-"
	database = "tasky"
)

func (m *MongoRepository) GetCollection(collection string) (*mongo.Collection, *model.Error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@testgo.6itcv.mongodb.net/%s?retryWrites=true&w=majority", usr, pwd, database)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		tempErr := &model.Error{
			Message: err.Error(),
		}
		fmt.Println(&tempErr.Message)
		return nil, tempErr
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		tempErr := &model.Error{
			Message: err.Error(),
		}
		fmt.Println(&tempErr.Message)
		return nil, tempErr
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		tempErr := &model.Error{
			Message: err.Error(),
		}
		fmt.Println(&tempErr.Message)
		return nil, tempErr
	}
	fmt.Println("Databases: ", databases)

	return client.Database(database).Collection(collection), nil
}
