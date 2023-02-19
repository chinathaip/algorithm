package main

/*

1. Create an empty hash table freq to keep track of the frequency of each element in nums1.

2. Iterate over nums1 and count the frequency of each element in freq.

3. Create an empty slice result to store the intersection of nums1 and nums2.

4. Iterate over nums2 and check if each element is in freq. If it is, add it to result and decrement its frequency in freq.

5. Return result.


Time complexity = O(m+n) -> m= len nums1, n= len nums2
- each element is processed atmost twice to fid intersection(once in each loop)

Space complexity = O(m+n)
- use map to store every element in nums1 = O(m)
- the intersect array could take m+n space in the worst case (every number is the same in both array)
*/

func intersect(nums1 []int, nums2 []int) []int {

	freq := map[int]int{}

	for _, num := range nums1 {
		freq[num] += 1
	}

	inter := []int{}

	for _, num := range nums2 {
		if freq[num] > 0 {
			inter = append(inter, num)
			freq[num] -= 1
		}
	}

	return inter
}
