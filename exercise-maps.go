package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	resMap := make(map[string]int)
	for _, v := range strings.Fields(s) {
		resMap[v]++
	}
	return resMap
}

func main() {
	wc.Test(WordCount)
}
