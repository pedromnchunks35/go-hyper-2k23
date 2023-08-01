package test

import (
	"context"
	"strings"
	t2 "tasks/protofiles"
	"testing"
)

func Test_Create_Task_Invalid(t *testing.T) {
	task := &t2.Task{}
	task.Description = ""
	_, err := Client.CreateTask(context.Background(), task)
	if err == nil || !strings.Contains(err.Error(), "please provide a valid task for saving") {
		t.Fatalf("must throw a error and container the given message")
	}
	task.Description = "Something"
	task.EndDate = ""
	_, err = Client.CreateTask(context.Background(), task)
	if err == nil || !strings.Contains(err.Error(), "please provide a valid task for saving") {
		t.Fatalf("must throw a error and container the given message")
	}
	task.EndDate = "2023/07/12 09:00:00"
	task.StartDate = ""
	_, err = Client.CreateTask(context.Background(), task)
	if err == nil || !strings.Contains(err.Error(), "please provide a valid task for saving") {
		t.Fatalf("must throw a error and container the given message")
	}
}

func Test_Create_Task(t *testing.T) {
	task := &t2.Task{}
	task.Description = "Need to do coffe"
	task.EndDate = "2023/07/12 09:00:00"
	task.StartDate = "2023/07/12 09:00:00"
	result, err := Client.CreateTask(context.Background(), task)
	if err != nil {
		t.Fatalf("it cannot throw a error, the task is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The received task was saved with id:") {
		t.Fatalf("must throw a success message")
	}
}

func Test_Update_Invalid(t *testing.T) {
	task := &t2.TaskMetaData{}
	task.IdTask = 100
	_, err := Client.TaskCompleted(context.Background(), task)
	if err == nil || !strings.Contains(err.Error(), "please provide a valid id") {
		t.Fatalf("it must throw a error saying that the id is invalid")
	}
	task.IdTask = -1
	_, err = Client.TaskCompleted(context.Background(), task)
	if err == nil || !strings.Contains(err.Error(), "cannot find the given id to update") {
		t.Fatalf("it must throw a error saying that it does not find the id")
	}
}

func Test_Update(t *testing.T) {
	task := &t2.Task{}
	task.Description = "Need to do coffe"
	task.EndDate = "2023/07/12 09:00:00"
	task.StartDate = "2023/07/12 09:00:00"
	result, err := Client.CreateTask(context.Background(), task)
	if err != nil {
		t.Fatalf("it cannot throw a error, the task is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The received task was saved with id:") {
		t.Fatalf("must throw a success message")
	}
	taskMeta := &t2.TaskMetaData{}
	taskMeta.IdTask = 1
	result, err = Client.TaskCompleted(context.Background(), taskMeta)
	if err != nil {
		t.Fatalf("it should not throw a error %v", err)
	}
	if !strings.Contains(result.Msg, "The task with id 1, got updated successfully") {
		t.Fatalf("it should throw a success message")
	}
}

func Test_Get_Tasks_Invalid(t *testing.T) {
	req := &t2.RequestMetaData{}
	req.OffSet = 1000
	req.PageSize = 2000
	_, err := Client.GetTasks(context.Background(), req)
	if err == nil || !strings.Contains(err.Error(), "the given metadata is invalid") {
		t.Fatalf("it should throw a error with metadata invalid")
	}
}

func Test_Get_Tasks(t *testing.T) {
	task := &t2.Task{}
	task.Description = "Need to do coffe"
	task.EndDate = "2023/07/12 09:00:00"
	task.StartDate = "2023/07/12 09:00:00"
	result, err := Client.CreateTask(context.Background(), task)
	if err != nil {
		t.Fatalf("it cannot throw a error, the task is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The received task was saved with id:") {
		t.Fatalf("must throw a success message")
	}

	task2 := &t2.Task{}
	task2.Description = "Need to do coffe v2"
	task2.EndDate = "2023/07/12 10:00:00"
	task2.StartDate = "2023/07/12 11:00:00"
	result, err = Client.CreateTask(context.Background(), task2)
	if err != nil {
		t.Fatalf("it cannot throw a error, the task is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The received task was saved with id:") {
		t.Fatalf("must throw a success message")
	}

	task3 := &t2.Task{}
	task3.Description = "Need to do coffe v3"
	task3.EndDate = "2023/07/13 10:00:00"
	task3.StartDate = "2023/07/14 11:00:00"
	result, err = Client.CreateTask(context.Background(), task3)
	if err != nil {
		t.Fatalf("it cannot throw a error, the task is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The received task was saved with id:") {
		t.Fatalf("must throw a success message")
	}

	req := &t2.RequestMetaData{}
	req.OffSet = 2
	req.PageSize = 3
	resultList, err := Client.GetTasks(context.Background(), req)
	if err != nil {
		t.Fatalf("it should not throw a error %v", err)
	}
	if len(resultList.TaskData) != 3 {
		t.Fatalf("it should have 3 items")
	}
	if resultList.TaskData[0].Id != 3 ||
		!strings.Contains(resultList.TaskData[0].Description, "Need to do coffe") ||
		!strings.Contains(resultList.TaskData[0].StartDate, "2023/07/12 09:00:00") ||
		!strings.Contains(resultList.TaskData[0].EndDate, "2023/07/12 09:00:00") ||
		resultList.TaskData[0].IsCompleted {
		t.Fatalf("it should have the data of the first task1")
	}
	if resultList.TaskData[1].Id != 4 ||
		!strings.Contains(resultList.TaskData[1].Description, "Need to do coffe v2") ||
		!strings.Contains(resultList.TaskData[1].StartDate, "2023/07/12 11:00:00") ||
		!strings.Contains(resultList.TaskData[1].EndDate, "2023/07/12 10:00:00") ||
		resultList.TaskData[1].IsCompleted {
		t.Fatalf("it should have the data of the first task2")
	}
	if resultList.TaskData[2].Id != 5 ||
		!strings.Contains(resultList.TaskData[2].Description, "Need to do coffe v3") ||
		!strings.Contains(resultList.TaskData[2].StartDate, "2023/07/14 11:00:00") ||
		!strings.Contains(resultList.TaskData[2].EndDate, "2023/07/13 10:00:00") ||
		resultList.TaskData[2].IsCompleted {
		t.Fatalf("it should have the data of the first task3")
	}
}
