package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCounter := make(map[string]int)
	for _, word := range words {
		wordCounter[word]++
	}
	return wordCounter
}

func main() {
	wc.Test(WordCount)
}
