package main

import (
	"fmt"
	"time"

	"tasky/database/connection"
	"tasky/model"
	"tasky/service/task"

	"github.com/spf13/viper"
)

func createTask(taskService task.TaskService) {
	task := model.Task{
		Description: "This is another test of a Task",
		UserId:      "zy12xw34vu56ts78",
		FinishAt:    time.Now().AddDate(0, 0, 8),
		IsDeleted:   false,
		IsMarked:    false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// task := model.Task{
	// 	Description: "This is a test of a Task",
	// 	UserId:      "ab12cd34ef56tgh78",
	// 	FinishAt:    time.Now().AddDate(0, 0, 3),
	// 	IsDeleted:   false,
	// 	IsMarked:    false,
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }
	taskService.Create(task)
}

func main() {
	fmt.Println("Hello from Tasky!")

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	mongodb := connection.New()
	taskService := task.New(mongodb)

	tasks, _ := taskService.Read()
	fmt.Println(*tasks)

	// createTask(*taskService)
	// taskToUpdate := model.Task{
	// 	Description: "This is a update of a Task",
	// 	UserId:      "zy12xw34vu56ts78",
	// 	FinishAt:    time.Now().AddDate(0, 0, 5),
	// 	IsDeleted:   false,
	// 	IsMarked:    false,
	// 	UpdatedAt:   time.Now(),
	// }

	// taskService.Update(taskToUpdate, "607b838fdf39612c80554a6a")

	// taskService.Delete("607b814fa091788c80ba27f9")
}
