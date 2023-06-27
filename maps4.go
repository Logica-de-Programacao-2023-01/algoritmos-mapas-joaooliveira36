package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagramGroups(words []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)

		if _, ok := groups[sortedWord]; ok {
			groups[sortedWord] = append(groups[sortedWord], word)
		} else {
			groups[sortedWord] = []string{word}
		}
	}

	return groups
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	words := []string{"listen", "silent", "elbow", "below", "state", "taste", "plate"}

	anagramMap := anagramGroups(words)

	for group, wordsInGroup := range anagramMap {
		fmt.Printf("%s: %v\n", group, wordsInGroup)
	}
}
