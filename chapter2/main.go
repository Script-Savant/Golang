package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g en")
	// lang := flag.StringVar("lang", "en", "The required language, e.g en")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// langauge represents the languahe's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "םלוע םולש",         // Hebrew
	"ur": "ایند ولیہ",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in the specified language
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
