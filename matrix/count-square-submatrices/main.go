package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}

	fmt.Println("matrix =", matrix)
	sum := countSquares(matrix)
	fmt.Println("sum = ", sum)

}

func countSquares(matrix [][]int) int {
	sum := 0

	// m = rows
	m := len(matrix)
	fmt.Println("m =", m)

	// n = columns
	n := len(matrix[0])
	fmt.Println("n =", n)

	// create dp matrix to store solutions to subproblems
	dpMatrix := make([][]int, m)

	for i := range dpMatrix {
		dpMatrix[i] = make([]int, n)
	}

	// iterate over matrix and compute subproblems
	// i = row iterator
	// j = column iterator
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 {
				// top row and left column
				dpMatrix[i][j] = matrix[i][j]
			} else {
				// all other matrix elements
				if matrix[i][j] == 1 {
					left := dpMatrix[i][j-1]
					diagonal := dpMatrix[i-1][j-1]
					top := dpMatrix[i-1][j]
					min := Min(left, Min(diagonal, top))
					dpMatrix[i][j] = min + 1
				}
			}

		}
	}
	fmt.Println("dpMatrix =", dpMatrix)

	for i := range dpMatrix {
		for j := range dpMatrix[i] {
			sum += dpMatrix[i][j]
		}

	}

	return sum
}

// Min returns the smaller int of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
