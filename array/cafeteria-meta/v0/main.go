package main

import (
	"fmt"
)

func main() {

	N := int64(15)

	K := int64(2)

	M := int32(3)

	S := []int64{11, 6, 14}

	r := getMaxAdditionalDinersCount(N, K, M, S)
	fmt.Println("r =", r)
}

/*
N = total number of seats at the table
K = number of seats to right and left of a diner that are closed
M = total number of diners currently seated at the table
S = slice representing the seat numbers for the seated diners
*/
func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {

	// additional diners
	count := int64(0)

	// base case, all chairs filled by diners
	if M == int32(N) {
		return count
	}

	// create hash table
	t := createTable(S, K, N)

	// loop from 1 to N
	for i := int64(1); i <= N; i++ {
		if t[i] == "person" {
			i += K
			continue
		} else if t[i] == "closed" {
			continue
		} else {
			// check social distancing
			isFree := true

			for j := int64(1); j <= K; j++ {
				if t[i+j] == "person" {
					isFree = false
				} else {
					t[i+j] = "closed"
				}
			}

			if isFree {
				count += 1
				t[i] = "person"
				i += K
			}
		}
	}
	fmt.Println("count =", count)
	return count
}

// CreateTable loops through the seated diners'
// slice, S, and creates a hash table to identify
// whether a seat is occupied or closed based on
// the social distancing guideline, K.
func createTable(S []int64, K int64, N int64) map[int64]string {
	t := make(map[int64]string)

	for i := 0; i < len(S); i++ {
		// store seated persons
		t[S[i]] = "person"

		// store closed seats
		for j := int64(1); j <= K; j++ {
			if S[i]+j <= N {
				t[S[i]+j] = "closed"
			}
			if S[i]-j >= 1 {
				t[S[i]-j] = "closed"
			}
		}
	}
	return t
}
