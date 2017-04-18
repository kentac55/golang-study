package main

import (
	"fmt"
	"math"
)

func calcNextNum(input float64, target float64) float64 {
	return input - (input*input-target)/(input*2)
}

func isFinished(curNum float64, a float64, b float64) bool {
	if curNum == a || curNum == b {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {

	var (
		nextNum = 1.0
		a       = -1.0
		b       = -1.0
	)

	for i := 1; !isFinished(nextNum, a, b); i++ {
		b = a
		a = nextNum
		nextNum = calcNextNum(nextNum, x)
	}
	return nextNum
}

func main() {
	const num = 2.0
	fmt.Println("自作関数の答え: ", Sqrt(num))
	fmt.Println("組込関数の答え: ", math.Sqrt(num))
}
