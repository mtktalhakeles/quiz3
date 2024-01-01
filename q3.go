package main

import (
	"fmt"
)

func findMostRepeated(data []string) string {
	// Create a map to store the frequency of each element.
	frequency := make(map[string]int)

	// Iterate over the array and count the occurrences of each element.
	for _, item := range data {
		frequency[item]++
	}

	// Find the element with the maximum frequency.
	var mostRepeated string
	maxFrequency := 0

	for item, freq := range frequency {
		if freq > maxFrequency {
			maxFrequency = freq
			mostRepeated = item
		}
	}

	return mostRepeated
}

func main() {

	input := []string{"apple", "pie", "apple", "red", "red", "red"}
	result := findMostRepeated(input)

	fmt.Println(result)
}
