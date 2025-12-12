package main

import "fmt"

const (
 	spanish = "Spanish"
 	french = "French"
 	norwegian = "Norwegian"

 	englishHelloPrefix = "Hello, "
 	spanishHelloPrefix = "Hola, "
 	frenchHelloPrefix = "Bonjour, "
 	norwegianHelloPrefix = "Hei, "
)
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {	
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case norwegian:
		prefix = norwegianHelloPrefix
	default:
		prefix = englishHelloPrefix 
	}
	return
}

func main() {

	fmt.Println(Hello("", ""))

}
