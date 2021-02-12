package workerpool2

import (
	"fmt"
)

var errorChanel = make(chan int, 1000)

// Task là 1 simple struct mà lưu giữ tất cả các thứ cần thiết để xử lý 1 task
// chúng ta truyền vào data mà hàm f cần xử lý
// hàm f nhận data rồi xử lý
type Task struct {
	Err 	error
	Data	interface{}
	f 		func(interface{}) error	
}

// NewTask tạo 1 task mới
func NewTask(f func(interface{}) error, data interface{}) *Task {
	return &Task{f: f, Data: data}
}

func process(workerID int, task *Task) {
	fmt.Printf("Worker %d processes task %v \n", workerID, task.Data)
	task.Err = task.f(task.Data)
	errorChanel <- task.Data.(int)
}