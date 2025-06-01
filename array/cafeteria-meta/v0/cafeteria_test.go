package main

import (
	"fmt"
	"testing"
	"time"
)

type Test struct {
	N        int64
	K        int64
	M        int32
	S        []int64
	expected int64
}

func TestFindAdditionalDinersCount(t *testing.T) {
	tests := []Test{
		{
			N:        12,
			K:        1,
			M:        4,
			S:        []int64{1, 4, 6, 10},
			expected: 2,
		},
		{
			N:        10,
			K:        1,
			M:        2,
			S:        []int64{2, 6},
			expected: 3,
		},
		{
			N:        15,
			K:        2,
			M:        3,
			S:        []int64{11, 6, 14},
			expected: 1,
		},
		{
			N:        16,
			K:        3,
			M:        2,
			S:        []int64{8, 4},
			expected: 2,
		},
		{
			N:        16,
			K:        3,
			M:        2,
			S:        []int64{16, 1},
			expected: 2,
		},
	}

	for idx, tst := range tests {
		fmt.Println()
		fmt.Println("Test", idx+1)
		fmt.Println("----------------------")
		fmt.Println("S =", tst.S)

		// start watch
		start := time.Now()
		r := getMaxAdditionalDinersCount(tst.N, tst.K, tst.M, tst.S)

		// calculate execution time
		duration := time.Since(start)
		fmt.Println("duration =", duration)

		if r != tst.expected {
			t.Fatalf("\nS: %v\nexpected %v, got %v", tst.S, tst.expected, r)
		}
	}
}
