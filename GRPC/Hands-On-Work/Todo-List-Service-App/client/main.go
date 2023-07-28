package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	t "tasks/protofiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InsertTask(client t.TaskManagerClient, newTask *t.Task) error {
	//? Create some context
	ctx := context.Background()
	//? Make a connection
	confirmation, err := client.CreateTask(ctx, newTask)
	if err != nil {
		return fmt.Errorf("some error occured when making the request: %v", err)
	}
	log.Printf("[Server] %v", confirmation)
	return nil
}

func UpdateTask(client t.TaskManagerClient, metadata *t.TaskMetaData) error {
	//? Create some context
	ctx := context.Background()
	//? Make the connection
	confirmation, err := client.TaskCompleted(ctx, metadata)
	if err != nil {
		return fmt.Errorf("some error occured when making the request: %v", err)
	}
	log.Printf("[Server] %v", confirmation)
	return nil
}

func GetTasks(client t.TaskManagerClient, metadata *t.RequestMetaData) error {
	//? Create some context
	ctx := context.Background()
	//? Make the connection
	taskList, err := client.GetTasks(ctx, metadata)
	if err != nil {
		return fmt.Errorf("some error occured when making the request: %v", err)
	}
	log.Printf("[Server] Requested data: \n")
	for _, task := range taskList.TaskData {
		log.Printf("%v\n", task)
	}
	return nil
}

var (
	new_task      = flag.Bool("newTask", false, "If we want or not a new task")
	description   = flag.String("description", "", "A description of a task")
	start_date    = flag.String("start_date", "", "Start date of a task")
	end_date      = flag.String("end_date", "", "End date of a task")
	complete_task = flag.Bool("completeTask", false, "If we want or not to make a task completed")
	off_set       = flag.Int("offSet", 0, "The offset")
	page_size     = flag.Int("pageSize", 0, "The pageSize")
	id            = flag.Int("id", 0, "The id of the task we want to update")
	get_tasks     = flag.Bool("getTasks", false, "If we want or not to retrieve all tasks")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:2000", opts...)
	if err != nil {
		log.Fatalf("something went wrong with the creation of the connection %v", err)
	}
	defer conn.Close()
	client := t.NewTaskManagerClient(conn)
	if *new_task {
		newTask := &t.Task{}
		newTask.Description = *description
		newTask.EndDate = *end_date
		newTask.StartDate = *start_date
		newTask.IsCompleted = false
		err = InsertTask(client, newTask)
		if err != nil {
			log.Fatalf("[Server] %v", err)
		}
	}
	if *complete_task {
		metadataId := &t.TaskMetaData{}
		metadataId.IdTask = int32(*id)
		err = UpdateTask(client, metadataId)
		if err != nil {
			log.Fatalf("[Server] %v", err)
		}
	}
	if *get_tasks {
		metadata := &t.RequestMetaData{}
		metadata.OffSet = int32(*off_set)
		metadata.PageSize = int32(*page_size)
		err = GetTasks(client, metadata)
		if err != nil {
			log.Fatalf("[Server] %v", err)
		}
	}

}
