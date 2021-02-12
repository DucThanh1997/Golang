package main

import (
	"fmt"
	"time"

	"worker_pool/workerpool2"
)

func main() {
	var allTask []*workerpool2.Task
	for i := 1; i <= 100; i++ {
		task := workerpool2.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool2.NewPool(allTask, 5)
	pool.Run()
}