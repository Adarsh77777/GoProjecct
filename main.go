package main

import (
	"fmt"
	"github.com/Adarsh77777/GoProjecct/manager"
	"github.com/Adarsh77777/GoProjecct/node"
	job "github.com/Adarsh77777/GoProjecct/task"
	"github.com/Adarsh77777/GoProjecct/worker"
	"github.com/docker/docker/client"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"os"
	"time"
)

func main() {
	t := job.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  job.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}
	te := job.TaskEvent{
		ID:        uuid.New(),
		State:     job.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}
	fmt.Printf("task: %v\n", te)
	fmt.Printf("task event: %v\n", te)

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]job.Task),
	}
	fmt.Printf("worker: %v\n", w)
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string][]job.Task),
		EventDb: make(map[string][]job.TaskEvent),
		Workers: []string{w.Name},
	}

	fmt.Printf("manager: %v\n", m)
	m.SelectWorker()
	m.SelectTasks()
	m.UpdateTasks()
	m.SendWork()

	n := node.Node{
		Name: "Node-1",
		Ip:   "192.168.1.1",
		// Cores:  4,
		Memory: 1024,
		Disk:   25,
		// Role:   "worker",
	}

	fmt.Printf("node: %v\n", n)

	fmt.Printf("create a test container\n")
	dockerTask, createResult := createContainer()
	if createResult.Error != nil {
		fmt.Print(createResult.Error)
		os.Exit(1)
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("stopping container %s\n", createResult.ContainerId)
	_ = stopContainer(dockerTask)
}
func createContainer() (*job.Docker, *job.DockerResult) {
	c := job.Config{
		Name:  "test-container-1",
		Image: "postgres:13",
		Env: []string{
			"POSTGRES_USER=cube",
			"POSTGRES_PASSWORD=secret",
		},
	}
	dc, _ := client.NewClientWithOpts(client.FromEnv)
	d := job.Docker{
		Client: dc,
		Config: c}
	result := d.Run()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil, nil
	}
	fmt.Printf(
		"Container %s is running with config %v\n", result.ContainerId, c)
	return &d, &result
}

func stopContainer(d *job.Docker) *job.DockerResult {
	result := d.Stop()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		return nil
	}
	fmt.Printf(
		"Container %s has been stopped and removed\n", result.ContainerId)
	return &result
}
