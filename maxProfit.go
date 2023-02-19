package main

/*


[7,1,5,3,6,4]
*/

func maxProfit(prices []int) int {
	globalMax := 0
	for i := 0; i < len(prices); i++ {
		localMax := 0
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > localMax {
				localMax = profit
			}
		}
		if localMax > globalMax {
			globalMax = localMax
		}
	}
	return globalMax
}
