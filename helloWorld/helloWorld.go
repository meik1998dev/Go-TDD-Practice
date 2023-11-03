package main

import "fmt"

func main() {
	fmt.Println(Hello("world", "Spanish"))
}

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	french             = "French"
	spanish            = "Spanish"
)

func Hello(name, lang string) string {

	if name == "" {
		name = "World"
	}

	return greetingPre(lang) + name
}

func greetingPre(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
