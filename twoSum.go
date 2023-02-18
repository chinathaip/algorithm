package main

/*
1. nested loop through the slice
2. see if the current num + the rest of num matches with target or not
3. yes -> return index of current num and that number
*/
func twoSum(nums []int, target int) []int {
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr = append(arr, i, j)
			}
		}
	}
	return arr
}

//Time complexity: O(n^2)
//Space complextiy: O(1)
