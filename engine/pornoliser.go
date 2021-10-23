package engine

import (
	"math/rand"
	"strings"
	"time"
)

func Pornolise(text string, swearLevel int, language string, entropy int) string {
	return proseConverter(nameConverter(text, swearLevel, language, entropy), swearLevel, language, entropy)
}

func nameConverter(text string, swearibility int, language string, entropy int) string {

	names := languageNames[language]
	match := 0
	returnText := ""
	textArray := strings.Fields(text)
	rand.Seed(int64(entropy) + time.Now().UnixNano())
	for _, word := range textArray {
		if word == strings.Title(word) {
			match = match + 1
		} else {
			match = 0
		}

		if match == 2 {
			if rand.Intn(100) <= swearibility { //len(names)
				returnText = returnText + " \"" + names[rand.Intn(len(names))] + "\""
			}
			match = 0
		}
		returnText = returnText + " " + word
	}

	return strings.TrimSpace(returnText)
}

func proseConverter(text string, swearibility int, language string, entropy int) string {

	dict := languageAdjectives[language]
	returnText := ""
	textArray := strings.Fields(text)
	suffix := ""
	rand.Seed(int64(entropy) + time.Now().UnixNano())
	for _, word := range textArray {
		if strings.HasSuffix(word, "ing") || strings.HasSuffix(word, "ed") || strings.HasSuffix(word, "s") {
			if rand.Intn(100) <= swearibility { //rand.Intn(len(dict))

				if strings.HasSuffix(word, "ing") {
					suffix = "ing"
				}
				if strings.HasSuffix(word, "ed") {
					suffix = "ed"
				}
				if strings.HasSuffix(word, "s") {
					suffix = "s"
				}

				word = dict[rand.Intn(len(dict))] + suffix
				suffix = ""
			}
		}
		returnText = returnText + " " + word
	}
	return strings.TrimSpace(returnText)
}
