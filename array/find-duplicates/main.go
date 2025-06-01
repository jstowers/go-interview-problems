package main

import "fmt"

func findDuplicates(arr1 []int, arr2 []int) []int {
	// define result slice, r
	var r []int

	fmt.Println("arr1 =", arr1)
	fmt.Println("arr2 =", arr2)

	// define map, m
	// the zero value of a map is nil
	// a nil map has no keys, nor can keys be added
	var m map[int]int

	// make returns a map of the given type, initialized
	// and ready to use
	m = make(map[int]int)

	// fill map with arr2 values
	for i := 0; i < len(arr2); i++ {
		m[arr2[i]] = 1
	}

	// loop over arr1 and check for duplicates in m
	for j := 0; j < len(arr1); j++ {
		if m[arr1[j]] == 1 {
			r = append(r, arr1[j])
		}
	}
	return r
}

func main() {
	arr1 := []int{1, 12, 15, 19, 20, 21}
	arr2 := []int{2, 15, 17, 19, 21, 25, 27}

	r := findDuplicates(arr1, arr2)
	fmt.Println("common elements =", r)
}
