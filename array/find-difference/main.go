package main

import "fmt"

/*
FindDifference returns a 2D array.  The first inner array is
a list of distinct integers in nums1 that are not present
in nums2.  The second inner array is a list of distinct integers
in nums2 that are not present in nums1.
*/
func findDifference(nums1 []int, nums2 []int) [][]int {

	r := [][]int{{}, {}}

	// create a hash map of values for nums1
	var m1 map[int]bool
	m1 = make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		if m1[nums1[i]] != true {
			m1[nums1[i]] = true
		}
	}

	// create a hash map of values for nums2
	var m2 map[int]bool
	m2 = make(map[int]bool)

	for j := 0; j < len(nums2); j++ {
		if m2[nums2[j]] != true {
			m2[nums2[j]] = true
		}
	}

	// loop through map1 and check map2 for the value
	// if it does not exist, add to result1 slice.
	for k := range m1 {
		if m2[k] != true {
			r[0] = append(r[0], k)
		}
	}

	// loop through map2 and check map1 for the value.
	// if it does not exist, add to result2 slice.
	for k := range m2 {
		if m1[k] != true {
			r[1] = append(r[1], k)
		}
	}

	// r = [r0, r1]
	return r
}

func main() {
	nums1 := []int{1, 2, 3, 3, 4, 7, 8}
	nums2 := []int{4, 5, 6, 8}

	result := findDifference(nums1, nums2)
	fmt.Println("result =", result)
}
