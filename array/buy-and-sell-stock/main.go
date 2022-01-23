package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	//prices := []int{7, 1, 5, 3, 6, 4}

	// below declaration imports a large data set for evaluation
	// prices := data

	profit := maxProfit(prices)
	fmt.Println("profit = ", profit)

}

// Solution using a single for loop
// Time complexity: O(n) => linear
// Takes time proportional to the input size
// Space complexity: O(1) => constant
// the program stores the three sames variables (buyPrice, profit, and maxProfit),
// regardless of the size of the prices[] dataset.

// https://www.baeldung.com/cs/space-complexity

func maxProfit(prices []int) int {

	buyPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - buyPrice

		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else if profit > maxProfit {
			maxProfit = profit
		}

		fmt.Println("i = ", i, "prices[i] =", prices[i], "buyPrice = ", buyPrice, "profit = ", profit, "maxProfit = ", maxProfit)
	}

	return maxProfit
}

// Solution using double for loop
// Time complexity: O(n^2)
/*
func maxProfit(prices []int) int {
	currProfit, max := 0, 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			fmt.Println("buyPrice =", prices[i], "sellPrice =", prices[j])
			currProfit = prices[j] - prices[i]
			if currProfit > max {
				max = currProfit
			}
		}
	}
	// if no profit, return 0
	if max <= 0 {
		return 0
	}
	return max
}
*/
