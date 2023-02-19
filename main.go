package main

import "fmt"

func main() {
	// x := containsDuplicate([]int{1, 2, 3, 4})
	// x := containsDuplicate([]int{1, 2, 3, 1})
	// x := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})

	// x := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

	// x := twoSum([]int{2, 7, 11, 15}, 9)
	// fmt.Println(x)

	// merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	// merge([]int{1, 0}, 1, []int{2}, 1)

	// x := intersect([]int{1, 2, 2, 1}, []int{2, 2})

	// x := maxProfit([]int{7, 1, 5, 3, 6, 4})

	x := matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4)
	fmt.Println(x)
}
