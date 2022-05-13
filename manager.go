package manager

import (
	//job "Project/task"
	job "github.com/Adarsh77777/GoProjecct/task"
	//job "Project/task"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pendning      queue.Queue
	TaskDb        map[string][]job.Task
	EventDb       map[string][]job.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
	Pending       queue.Queue
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}
func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}
func (m *Manager) SelectTasks() {
	fmt.Println("I will select tasks")
}
func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
