package main

import "fmt"

func sum(nums []int, c chan int) {
	sum := 0
	for num := range nums {
		sum += num
	}
	c <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)
	s1, s2 := <-c, <-c
	fmt.Printf("sum := %v", s1+s2)

}
