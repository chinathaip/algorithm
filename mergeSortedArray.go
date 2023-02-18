package main

/*
The Merge Sorted Array problem asks you to merge two sorted arrays nums1 and nums2 into nums1 such that the resulting array is also sorted.

 1. initialize three pointers:
    i points to the last element in the nums1 array that is not part of the merged array
    j points to the last element in the nums2 array
    k points to the last position in nums1.

 2. iterate over the arrays using a while loop. While i and j are both greater than or equal to 0
    2.1 compare the elements at nums1[i] and nums2[j]
    2.2 insert the larger element at nums1[k].
    2.3 decrement i or j depending on which element was larger
    2.4 decrement k

 3. there may be some elements remaining in nums2 that have not been merged into nums1.
    3.1 use a second while loop that inserts any remaining elements in nums2 into nums1.

Finally, nums1 contains the merged and sorted array.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
