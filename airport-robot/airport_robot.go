package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

type German struct{}
type Italian struct{}
type Portuguese struct{}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(a string) string {
	return fmt.Sprintf("Hallo %v!", a)
}

func (g Italian) LanguageName() string {
	return "Italian"
}

func (g Italian) Greet(a string) string {
	return fmt.Sprintf("Ciao %v!", a)
}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}

func (g Portuguese) Greet(a string) string {
	return fmt.Sprintf("Olá %v!", a)
}
