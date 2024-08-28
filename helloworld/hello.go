package helloworld

func GreetingPrefix(language string) string {
	const spanish = "Spanish"
	const french = "French"

	const englishPrefix = "Hello, "
	const spanishPrefix = "Hola, "
	const frenchPrefix = "Bonjour, "

	switch language {
	case spanish:
		return spanishPrefix
	case french:
		return frenchPrefix
	default:
		return englishPrefix
	}
}

func Hello(name string, language string, punctuation string) string {
	switch language {
	case "English":
		return GreetingPrefix(language) + name + punctuation
	case "Spanish":
		return GreetingPrefix(language) + name + punctuation
	case "French":
		return GreetingPrefix(language) + name + punctuation
	default:
		return GreetingPrefix(language) + "world" + punctuation
	}
}
