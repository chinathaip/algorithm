package main

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := low + (high-low)/2
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			low = mid + 1 // skip mid since we already check
		} else {
			high = mid
		}
	}

	return -1
}
