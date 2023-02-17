package main

import "fmt"

func main() {
	// x := containsDuplicate([]int{1, 2, 3, 4})
	// x := containsDuplicate([]int{1, 2, 3, 1})
	// x := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})

	x := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(x)
}
