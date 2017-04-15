package main

import (
	"fmt"
	"math"
)

func calcNextNum(input float64, target float64) float64 {
	return input - (input*input-target)/(input*2)
}

func isFinished(current float64, past1stNum float64, past2ndNum float64) bool {
	if current == past1stNum || current == past2ndNum {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {

	if x == 0 {
		fmt.Println(0)
		return 0.0
	} else if x < 0 {
		fmt.Println("E: complex")
		return 0.0
	}

	var (
		nextNum      = 1.0
		past1stNum   = -1.0
		past2ndNum   = -1.0
		maxTryCounts = 100
	)

	for i := 1; !isFinished(nextNum, past1stNum, past2ndNum) && i <= maxTryCounts; i++ {
		past2ndNum = past1stNum
		past1stNum = nextNum
		nextNum = calcNextNum(nextNum, x)
		if i == maxTryCounts {
			fmt.Println("reached maxTryCounts.")
		}
	}
	return nextNum
}

func main() {
	num := 2.0
	fmt.Println("自作関数の答え: ", Sqrt(num))
	fmt.Println("組込関数の答え: ", math.Sqrt(num))
}
