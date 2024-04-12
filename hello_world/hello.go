package hello_world


const (
	englishLanguage = "english"
	spanishLanguage = "spanish"
	frenchLanguage  = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getGreetingPrefix(language) + name + "!"
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLanguage:
		prefix = frenchHelloPrefix
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case englishLanguage:
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

