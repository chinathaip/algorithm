package main

/*

1. loop through the slice
2. try to add the most sum
	2.1 see if the num


Time complexity: O(N)
Space complexity: O(1)
*/

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		//if nums[i] alone is greater than combined with currentSum, set value as nums[i]
		currentSum = max(nums[i], currentSum+nums[i])
		//check if currentSum is greater than the max sum we already have or not
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
