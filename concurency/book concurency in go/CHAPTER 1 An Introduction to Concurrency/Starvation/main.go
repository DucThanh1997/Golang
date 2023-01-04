package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var shareLock sync.Mutex
	const runtime = 1 * time.Second
	// cái greedy này chạy lấy hết time của cái polite dưới kia dẫn đến việc nó chạy được nhiều hơn thằng polite. 
	// Nó lock cả hàm còn thằng polite chỉ lock khi nó cần
	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			shareLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			shareLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			shareLock.Lock()
			time.Sleep(1*time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1*time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1*time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
