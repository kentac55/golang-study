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

func isFinished(curNum float64, a float64, b float64) bool {
	if curNum == a || curNum == b {
		return true
	}
	return false
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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
	return nextNum, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
