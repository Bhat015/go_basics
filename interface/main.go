package main

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	//logic for english
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	//logic for spanish
	return "Hola!"
}

func printGreeting(b bot) {
	println(b.getGreeting())
}
