package main

/*
1. loop through the slice
2. store frequency of each num in the slice in Map
3. see if any frequnecy is more than 1


*/

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	//put num in map
	for _, num := range nums {
		if _, found := m[num]; found {
			m[num] = m[num] + 1
		} else {
			m[num] = 1
		}
	}

	//check if freq more than 1
	for _, freq := range m {
		if freq > 1 {
			return true
		}
	}
	return false

	//Time complexity: O(N)
	//Space complexity: O(N)
}
