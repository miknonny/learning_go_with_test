package main

const (
	SpanishHelloPrefix = "Hola, "
	EnglishHelloPrefix = "Hello, "
	FrenchHelloPrefix  = "Bonjour, "
)

type Language int

const (
	English Language = iota
	Spanish
	French
)

func Hello(name string, lang Language) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang Language) string {
	var prefix string

	switch lang {
	case English:
		prefix = EnglishHelloPrefix
	case Spanish:
		prefix = SpanishHelloPrefix
	case French:
		prefix = FrenchHelloPrefix
	}

	return prefix
}
