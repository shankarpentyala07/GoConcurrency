/*
1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

*/

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t, sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		walkRecursive(t.Left, ch)
		ch <- t.Value
		walkRecursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 store the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2

		if val1 != val2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	// Test the Walk function
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	fmt.Println("Values from the channel:")
	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Println(val)
	}

	// Test the Same function
	fmt.Println("Test Same:")
	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(tree.New(1), tree.New(1)))   // should return true
	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))   // should return false
	fmt.Println("Same(tree.New(2), tree.New(2)):", Same(tree.New(2), tree.New(2)))   // should return true
	fmt.Println("Same(tree.New(2), tree.New(3)):", Same(tree.New(2), tree.New(3)))   // should return false
}
