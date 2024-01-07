package main

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	// get left-hand side product for each element
	product := 1
	left := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		left[i] = product
		product = product * nums[i]
	}

	// get right-hand side product for each element
	product = 1
	right := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = product
		product = product * nums[i]
	}

	// combine left and right together
	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}
