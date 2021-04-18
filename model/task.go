package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id          primitive.ObjectID `bson:"_id" json:ID,omitempty`
	Description string             `json:Description, omitempty`
	UserId      string             `json:userId, omitempty`
	FinishAt    time.Time          `bson:"finish_at" json:finish_at, omitempty`
	IsDeleted   bool               `json:isDeleted, omitempty`
	IsMarked    bool               `json:isMarked, omitempty`
	CreatedAt   time.Time          `bson:"created_at" json:created_at`
	UpdatedAt   time.Time          `bson:"updated_at" json:updated_at,omitempty`
}

type Tasks []Task
