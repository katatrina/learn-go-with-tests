package helloworld

import (
	"fmt"
)

const (
	vietnamese = "Vietnamese"
	french     = "French"
	
	englishHelloPrefix    = "Hello, "
	vietnameseHelloPrefix = "Xin chào, "
	frenchHelloPrefix     = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case vietnamese:
		prefix = vietnameseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
