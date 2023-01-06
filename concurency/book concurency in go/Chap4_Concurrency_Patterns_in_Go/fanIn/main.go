package main

import "fmt"

func printSth(msg string) chan string {
	result := make(chan string)

	go func() {
		for i := 0; i <= 5; i++ {
			result <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return result
}

func fanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case <- chan1:
				c <- <-chan1
			case <- chan2:
				c <- <-chan2
			}
		}
	}()
	return c
}

func main() {
	coffe := printSth("coffee")
	bread := printSth("bread order")
	serve := fanIn(coffe, bread)
	for i := 0; i <= 5; i++ {
		fmt.Println("receive from ", <- serve)
	}
	fmt.Println("finish")
}
