package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// note you don't have to have the short variable name defined
// (e.g. eb for englishBot)
func (englishBot) getGreeting() string {
	// "Very" custom logic for english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// "Very" custom logic for spanish greeting
	return "Hola!"
}
