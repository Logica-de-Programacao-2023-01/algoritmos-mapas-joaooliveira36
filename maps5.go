package main

import "fmt"

func characterFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)

	for _, char := range s {
		frequency[char]++
	}

	return frequency
}

func main() {
	str := "Hello, World!"
	frequencies := characterFrequency(str)

	for char, count := range frequencies {
		fmt.Printf("Character: %c, Frequency: %d\n", char, count)
	}
}
