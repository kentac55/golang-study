package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordsSlice := strings.Fields(s)
	resMap := make(map[string]int)
	for i, v := range wordsSlice {
		elem, ok := resMap[v]
		if ok {
			resMap[wordsSlice[i]] = elem + 1
		} else {
			resMap[wordsSlice[i]] = 1
		}
	}
	return resMap
}

func main() {
	wc.Test(WordCount)
}
