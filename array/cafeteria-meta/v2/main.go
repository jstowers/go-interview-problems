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

	r := getMaxAdditionalDinersCount(N, K, M, S)
	fmt.Println("r =", r)
}

func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {

	// additional diners
	count := int64(0)

	// minimum number of seats needed to add a diner
	// and satisfy social distancing requirements
	min := K + 1

	// sort S
	s := int64Slice(S)
	sort.Sort(s)

	// slices package not present in Meta
	//slices.Sort(S)
	//fmt.Println("S =", S)

	// additional diners before first seated diner
	count += (S[0] - 1) / min

	// additional diners after last seated diner
	count += (N - S[M-1]) / min

	// additional diners between segments of seated diners
	for i := int32(0); i < M-1; i++ {
		diner1 := S[i]
		diner2 := S[i+1]
		count += (diner2 - diner1 - min) / min
	}

	fmt.Println("count =", count)
	return count
}
