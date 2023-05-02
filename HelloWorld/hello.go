package HelloWorld

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// print hello, world
func Hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}

	return getPrefix(language) + name

}

// getting prefix
func getPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix

	}
	return
}
