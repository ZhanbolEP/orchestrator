package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	"Orchestrator/task"
)

type Worker struct {
	Name		string
	queue		queue.Queue
	Db			map[uuid.UUID]task.Task
	TaskCount	int
}

func (w *Worker) CollectStats() {
	fmt.Println("Will collect stats")
}

func (w *Worker) RunTask(){
	fmt.Println("Start/Stop task")
}

func (w *Worker) StartTask(){
	fmt.Println("Start task")
}

func (w *Worker) StopTask(){
	fmt.Println("Stop task")
}