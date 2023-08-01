package todoListServerImpl

import (
	"context"
	"fmt"
	"sync"
	t "tasks/protofiles"
)

type Task struct {
	*t.UnimplementedTaskManagerServer
	tasks        []*t.TaskData
	tasksMutex   sync.RWMutex
	counter      int
	counterMutex sync.RWMutex
}

// ? Function to init the TaskManager
func InitTaskManager() *Task {
	return &Task{tasks: []*t.TaskData{}, counter: 1}
}

// ? Function to create a task
func (taskManager *Task) CreateTask(ctx context.Context, newTask *t.Task) (*t.Confirmation, error) {
	//? Checking the data
	if newTask.Description == "" || newTask.EndDate == "" || newTask.StartDate == "" {
		return nil, fmt.Errorf("please provide a valid task for saving")
	}
	//? Block tasks
	taskManager.tasksMutex.Lock()
	taskManager.counterMutex.Lock()
	defer taskManager.tasksMutex.Unlock()
	defer taskManager.counterMutex.Unlock()
	//? Create new task
	taskToSave := &t.TaskData{}
	taskToSave.Description = newTask.Description
	taskToSave.StartDate = newTask.StartDate
	taskToSave.EndDate = newTask.EndDate
	taskToSave.IsCompleted = newTask.IsCompleted
	taskToSave.Id = int32(taskManager.counter)
	taskManager.tasks = append(taskManager.tasks, taskToSave)
	//? Create confirmation
	confirmation := &t.Confirmation{}
	confirmation.Msg = fmt.Sprintf("The received task was saved with id: %v", taskToSave.Id)
	taskManager.counter++
	return confirmation, nil
}

// ? Function to update a task to completed
func (taskManager *Task) TaskCompleted(ctx context.Context, metadata *t.TaskMetaData) (*t.Confirmation, error) {
	taskManager.tasksMutex.Lock()
	defer taskManager.tasksMutex.Unlock()
	taskManager.counterMutex.Lock()
	defer taskManager.counterMutex.Unlock()
	if int(metadata.IdTask) > taskManager.counter {
		return nil, fmt.Errorf("please provide a valid id")
	}
	for _, task := range taskManager.tasks {
		if task.Id == metadata.IdTask {
			task.IsCompleted = true
			confirmation := &t.Confirmation{}
			confirmation.Msg = fmt.Sprintf("The task with id %v, got updated successfully", metadata.IdTask)
			return confirmation, nil
		}
	}
	return nil, fmt.Errorf("cannot find the given id to update")
}

// ? Function to retrieve tasks
func (taskManager *Task) GetTasks(ctx context.Context, metadata *t.RequestMetaData) (*t.TaskList, error) {
	taskManager.tasksMutex.Lock()
	defer taskManager.tasksMutex.Unlock()
	list := &t.TaskList{}
	//? Make the checks
	if int(metadata.OffSet) > len(taskManager.tasks) || (int(metadata.OffSet)+int(metadata.PageSize)) > len(taskManager.tasks) {
		return nil, fmt.Errorf("the given metadata is invalid")
	}
	//? Get the data
	list.TaskData = taskManager.tasks[int(metadata.OffSet):(int(metadata.OffSet) + int(metadata.PageSize))]
	return list, nil
}
