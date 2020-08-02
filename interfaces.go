package main

import "fmt"

type bot interface {
	//getGreeting(string, int) (string, error)  (args) (output)
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func showInterface() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
