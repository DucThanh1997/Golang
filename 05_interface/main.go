package main

import "fmt"

// có nghĩa là thằng nào có cái chức năng get greeting thì thuộc dạng bot này 
// và được sử dụng những cái của bot
type bot interface {
	getGreeting() string
}


type englishBot struct{}
type spanishBot struct{}

func (e englishBot) getGreeting() string {
	return "hello"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (s spanishBot) getGreeting() string {
	return "hola"
}

func printGreetingS(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	fmt.Println(sb)
	printGreeting(eb)
}
