package main

/*
 1. loop through the slice
 2. get the complement by doing target - current num
 3. check if the map already contains the complement -->
    3.1 if yes -> return the index of the complement and the index of current num
    3.2 if not contain -> add to complement to the map along with its index

[2,7,11,15]

9-2 = 7
9-7 = 2
9-11 = -2
9-15 = -6

	{
		7: 0
		2: 1
		-2:2
		-6:3
	}
*/
func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num
		if complementIndex, found := numToIndex[complement]; found {
			return []int{complementIndex, i}
		}
		numToIndex[num] = i
	}
	return nil
}

//Time complexity: O(n)
//Space complextiy: O(n)
