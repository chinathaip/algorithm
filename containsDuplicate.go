package main

/*
1. loop through the slice
2. store each num into map, set value as true
3. if it's already in map, return true


*/

func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
