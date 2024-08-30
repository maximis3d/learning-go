package main

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}
	return gettingPrefix(language) + name

}

func gettingPrefix(language string) (prefix string) {
	switch language {
	default:
		prefix = englishHelloPrefix

	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return
}
