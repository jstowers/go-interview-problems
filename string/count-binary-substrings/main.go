package main

import (
	"fmt"
	"math"
)

func main() {
	//bString := "10101"
	bString := "11000110"
	//bString := "00110011"
	//bString := "00110"
	count := countBinarySubstrings(bString)
	fmt.Println("substring count =", count)
}

func countBinarySubstrings(s string) int {

	// Check for an empty string
	if len(s) == 0 {
		return 0
	}

	// Initialize groups slice with 1 character
	groups := []int{1}

	// Loop through the string and group equal numbers (0's and 1's) into a groups[]
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			groups = append(groups, 1)
		} else {
			groups[len(groups)-1]++
		}
	}
	fmt.Println("groups = ", groups)

	// Compare consecutive group elements and choose the smallest number.
	// This number represents the number of substrings that can be created
	// between the two group elements.
	// Keep a running sum and return at the end of the loop
	sum := 0
	for i := 1; i < len(groups); i++ {
		minValue := math.Min(float64(groups[i]), float64(groups[i-1]))
		sum += int(minValue)
	}
	return sum
}
