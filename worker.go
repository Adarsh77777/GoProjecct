package worker

import (
	"fmt"
	job "github.com/Adarsh77777/GoProjecct/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Queue     queue.Queue
	Db        map[uuid.UUID]job.Task
	TaskCount int
	Name      string
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}
func (w *Worker) RunTask() {
	fmt.Println("I will run a task")
}
func (w *Worker) StartTask() {
	fmt.Println("I will start a task")
}
func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}
