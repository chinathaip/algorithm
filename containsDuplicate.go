package main

import (
	"sort"
)

/*
1. sort the slice
2. check if the current value is equal to the previous value or not
3. if equal -> return true

[1,2,3,1] --> [1,1,2,3]

This solution takes more time, but less space

*/

func containsDuplicate(nums []int) bool {
	sort.Ints(nums) //sort take O(n log n)

	for index, _ := range nums {
		if index == 0 {
			continue
		}

		if nums[index] == nums[index-1] {
			return true
		}
	}
	return false

	//Time complexity: O(n log n)
	//Space complexity: O(1)
}
