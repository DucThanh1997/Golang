package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case a, oke := <-c:
		fmt.Println("a: ", a)
		fmt.Println("oke: ", oke)
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
