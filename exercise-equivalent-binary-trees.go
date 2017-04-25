
package main

import "golang.org/x/tour/tree"

import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	if t.Left == nil && t.Right == nil {
		ch <- t.Value
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	isSame := true
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <- ch1 != <- ch2 {
			isSame = false
			break
		}
	}
	close(ch1)
	close(ch2)
	return isSame
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

	}()
	println(Same(tree.New(1), tree.New(1)))
	println(Same(tree.New(1), tree.New(2)))
}
