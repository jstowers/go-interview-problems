package main

import (
	"reflect"
	"testing"
)

func TestReorderLogFiles(t *testing.T) {

	// create a test structure of different log message scenarios
	type test struct {
		input []string
		want  []string
	}

	tests := []test{
		{
			input: []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			want:  []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			input: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			want:  []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
	}

	for _, test := range tests {
		output := reorderLogFiles(test.input)
		if !reflect.DeepEqual(output, test.want) {
			t.Fatalf("The reordering of input log is incorrect.\n Got:  %v\n Want: %v\n", output, test.want)
		}
	}
}
