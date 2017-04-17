package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	resMap := make(map[string]int)
	for _, v := range strings.Fields(s) {
		elem, ok := resMap[v]
		if ok {
			resMap[v] = elem + 1
		} else {
			resMap[v] = 1
		}
	}
	return resMap
}

func main() {
	wc.Test(WordCount)
}
