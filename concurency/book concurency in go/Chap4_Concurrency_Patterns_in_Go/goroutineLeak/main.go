package main

import (
	"fmt"
	"time"
)

// leak
// func main() {
// 	doWork := func(strings <-chan string) <-chan interface{} {
// 		completed := make(chan interface{})

// 		go func() {
// 			defer fmt.Println("doWork exited.")
// 			defer close(completed)
// 			for s := range strings {
// 				// Do something interesting
// 				fmt.Println(s)
// 			}
// 		}()
// 		return completed
// 	}

// 	doWork(nil)
// 	// Perhaps more work is done here
// 	fmt.Println("Done.")

// }

// solve
func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} { // số 1
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done: // số 2
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() { // số 3
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated // số 4
	fmt.Println("Done.")

	
}
