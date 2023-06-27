package main

import "fmt"

func countLetters(word string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, letter := range word {
		letterCount[letter]++
	}

	return letterCount
}

func wordLetterCountMap(sentence string) map[string]map[rune]int {
	wordLetterCount := make(map[string]map[rune]int)

	words := splitIntoWords(sentence)

	for _, word := range words {
		letterCount := countLetters(word)
		wordLetterCount[word] = letterCount
	}

	return wordLetterCount
}

func splitIntoWords(sentence string) []string {
	words := make([]string, 0)
	currentWord := ""

	for _, char := range sentence {
		if char != ' ' {
			currentWord += string(char)
		} else if currentWord != "" {
			words = append(words, currentWord)
			currentWord = ""
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}

func main() {
	sentence := "The quick brown fox jumps over the lazy dog"

	wordLetterCount := wordLetterCountMap(sentence)

	for word, letterCount := range wordLetterCount {
		fmt.Println("Word:", word)
		for letter, count := range letterCount {
			fmt.Printf("Letter: %c, Count: %d\n", letter, count)
		}
		fmt.Println()
	}
}
