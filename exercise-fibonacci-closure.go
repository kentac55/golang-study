package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pastA := 0
	pastB := 1
	return func() int {
		sum := pastA + pastB
		pastA = pastB
		pastB = sum
		return sum - pastA
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
