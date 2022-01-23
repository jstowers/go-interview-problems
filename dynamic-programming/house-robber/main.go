package main

import (
	"fmt"
)

func main() {

	// if 1 home => 1
	//nums := []int{100}

	// if 2 home => max(1, 2)
	//nums := []int{100, 200}

	// if 3 home => max (1 + 3, 2)
	// Expected: 7
	//nums := []int{4, 7, 2}

	// if 4 homes => max (1 + 3, 2 + 4, 1 + 4)
	// Expected: 4
	//nums := []int{2, 1, 1, 2}

	// if 5 homes => max (1 + 3 + 5, 2 + 4, 2 + 5)
	// Expected: 14  (2 + 5)
	//nums := []int{1, 10, 2, 3, 4}

	// Expected: 12  (1 + 3 + 5)
	// nums := []int{2, 7, 9, 3, 1}

	// Expected: 200 (2 + 5)
	nums := []int{1, 100, 3, 4, 100, 6}

	/*

		index = 0

			Pick max(nums[index], nums[index+1])

				// Compare 1 + 3 and 1 + 4
				If nums[0] => pick max(nums[index + 2], nums[index + 3])

				// Compare 2 + 4 and 2 + 5
				If nums[1] => pick max(nums[index + 1 + 2], nums[index + 1 +3])


	*/
	sum := rob(nums)
	fmt.Println("sum = ", sum)

}

// Recursive solution using dp[] to store the max amount that could be robbed
// at each step.
/*
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	// define solution matrix
	dp := make([]int, len(nums))

	// base cases
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])

	// solve dp subproblems
	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}
*/

// Recursive solution that only uses a sum.
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	previousSum := nums[0]
	currentSum := Max(nums[0], nums[1])

	// solve dp subproblems
	for i := 2; i < len(nums); i++ {
		previousSum = currentSum
		currentSum = Max(previousSum+nums[i], currentSum)
	}

	return currentSum
}

// Max returns the larger int of x or y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
