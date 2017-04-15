package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func calcNextNum(input float64, target float64) float64 {
	return input - (input*input-target)/(input*2)
}

func isFinished(current float64, past1stNum float64, past2ndNum float64) bool {
	if current == past1stNum || current == past2ndNum {
		return true
	}
	return false
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
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
	return nextNum, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
