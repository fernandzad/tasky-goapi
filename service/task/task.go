package task

import (
	"context"
	"fmt"
	"time"

	"tasky/database/connection"
	"tasky/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	mongodb connection.NewMongoRepository
}

type NewTaskService interface {
	Create(task model.Task) *model.Error
	Read() (*model.Tasks, *model.Error)
	Update(task model.Task, taskId string) *model.Error
	Delete(taskId string) *model.Error
}

func New(mongodb connection.NewMongoRepository) *TaskService {
	return &TaskService{mongodb}
}

func (ts *TaskService) Create(task model.Task) *model.Error {

	var collectionName string = "task"
	var collection, customErr = ts.mongodb.GetCollection(collectionName)
	defer ts.mongodb.Disconnect()
	if customErr != nil {
		return customErr
	}
	var ctx = context.Background()

	id := primitive.NewObjectID()
	task.Id = id

	_, err := collection.InsertOne(ctx, task)

	if err != nil {
		return &model.Error{
			Message: err.Error(),
		}
	}

	fmt.Println("Pokemon successfully created")
	return nil
}

func (ts *TaskService) Read() (*model.Tasks, *model.Error) {
	var tasks model.Tasks
	var collectionName string = "task"
	var collection, customErr = ts.mongodb.GetCollection(collectionName)
	defer ts.mongodb.Disconnect()
	if customErr != nil {
		return nil, customErr
	}
	var ctx = context.Background()

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, &model.Error{
			Message: err.Error(),
		}
	}

	for cur.Next(ctx) {
		var task model.Task
		err := cur.Decode(&task)

		if err != nil {
			return nil, &model.Error{
				Message: err.Error(),
			}
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func (ts *TaskService) Update(task model.Task, taskId string) *model.Error {
	var collectionName string = "task"
	var collection, customErr = ts.mongodb.GetCollection(collectionName)
	defer ts.mongodb.Disconnect()
	if customErr != nil {
		return customErr
	}
	var ctx = context.Background()

	oid, _ := primitive.ObjectIDFromHex(taskId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"description": task.Description,
			"userid":      task.UserId,
			"updated_at":  time.Now(),
			"finish_at":   task.FinishAt,
			"isdeleted":   task.IsDeleted,
			"ismarked":    task.IsMarked,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return &model.Error{
			Message: err.Error(),
		}
	}

	return nil
}

func (ts *TaskService) Delete(taskId string) *model.Error {
	var collectionName string = "task"
	var collection, customErr = ts.mongodb.GetCollection(collectionName)
	defer ts.mongodb.Disconnect()

	if customErr != nil {
		return customErr
	}
	var ctx = context.Background()

	oid, _ := primitive.ObjectIDFromHex(taskId)

	filter := bson.M{"_id": oid}

	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return &model.Error{
			Message: err.Error(),
		}
	}

	return nil
}
