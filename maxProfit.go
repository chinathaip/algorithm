package main

/*

Time Complexity = O(n) - loop through array once
Space Complexity = O(1) - no extra space
*/

func maxProfit(prices []int) int {
	bestPrice := 100_001
	maximumProfit := 0

	for _, price := range prices {

		//finding the lowest buying price
		if price < bestPrice {
			bestPrice = price
		}

		//find which gives the most profit
		profit := price - bestPrice
		if profit > maximumProfit {
			maximumProfit = profit
		}
	}
	return maximumProfit
}
