package main

import "fmt"

const englishPrefix = "Hello, "
const spanish = "Spanish"
const spanishPrefix = "Hola, "
const frensh = "Frensh"
const frenshPrefix = "Bonjour, "
const portuguese = "Portuguese"
const portuguesePrefix = "Olá, "
const japanese = "Japanese"
const japanesePrefix = "こんにちは, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		return spanishPrefix
	case frensh:
		return frenshPrefix
	case portuguese:
		return portuguesePrefix
	case japanese:
		return japanesePrefix
	default:
		return englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
