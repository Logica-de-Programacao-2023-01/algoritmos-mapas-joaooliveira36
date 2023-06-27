package main

import "fmt"

func mergeWordCountMaps(maps []map[string]int) map[string]int {
	result := make(map[string]int)

	for _, wordCountMap := range maps {
		for word, count := range wordCountMap {
			result[word] += count
		}
	}

	return result
}

func main() {
	text1 := map[string]int{"apple": 2, "banana": 3, "orange": 1}
	text2 := map[string]int{"banana": 1, "grape": 2, "kiwi": 3}
	text3 := map[string]int{"orange": 2, "kiwi": 1, "strawberry": 4}

	wordCountMaps := []map[string]int{text1, text2, text3}

	mergedMap := mergeWordCountMaps(wordCountMaps)

	for word, count := range mergedMap {
		fmt.Printf("Word: %s, Count: %d\n", word, count)
	}
}
