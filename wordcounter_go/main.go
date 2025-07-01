package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello, today we are learning go language and coding in go language"
	words := strings.Split(text, " ")

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	fmt.Println("Words count:")
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
