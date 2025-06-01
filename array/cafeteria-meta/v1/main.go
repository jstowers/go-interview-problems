package main

import (
	"fmt"
	"sort"
)

type int64Slice []int64

// define sort methods for sort.Sort interface
func (a int64Slice) Len() int {
	return len(a)
}

func (a int64Slice) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a int64Slice) Less(i int, j int) bool {
	return a[i] < a[j]
}

func (a int64Slice) String() (s string) {
	sep := "" // for printing separating commas
	for _, el := range a {
		s += sep
		sep = ", "
		s += fmt.Sprintf("%d", el)
	}
	return
}

func createSlice(N int64) []int64 {

	S := []int64{}

	for i := int64(2); i < N; i++ {
		S = append(S, i)
		i += 2 * i
	}

	return S
}

func main() {

	// sample #1
	// N := int64(10)
	// K := int64(1)
	// M := int32(2)
	// S := []int64{6, 2}

	// sample #2
	N := int64(15)
	K := int64(2)
	M := int32(3)
	S := []int64{11, 6, 14}

	// sample #3
	// N := int64(100)
	// K := int64(1)
	// S := createSlice(N)
	// M := int32(len(S))

	r := getMaxAdditionalDinersCount(N, K, M, S)
	fmt.Println("r =", r)
}

func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {

	// additional diners
	count := int64(0)

	// sort S
	// slices package not present in Meta
	//slices.Sort(S)
	//fmt.Println("S =", S)

	s := int64Slice(S)
	sort.Sort(s)

	// loop over S
	for i := 0; i <= len(S); i++ {

		// number of empty seats per segment
		empty := int64(0)

		// minimum number of seats needed to add a diner
		// and satisfy social distancing requirements
		min := int64(0)

		// process each table segment from start to end
		if i == 0 {
			// start: start of table to first diner
			empty = S[i] - 1
			min = K + 1
		} else if i <= len(S)-1 {
			// middle: each segment between seated diners
			empty = S[i] - S[i-1] - 1
			min = 2*K + 1
		} else {
			// end: last diner to end of table
			empty = N - S[i-1]
			min = K + 1
		}

		if empty >= min {
			maxPoss := int64(empty / min)
			count += maxPoss
		}
	}

	fmt.Println("count =", count)
	return count
}
