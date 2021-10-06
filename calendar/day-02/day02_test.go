package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Test that replacing a value in the array works
func TestReplaceValueInArray(t *testing.T) {
	var tests = []struct {
		value, postition int
		inputArray       []int
		want             []int
	}{
		{12, 1, []int{1, 0, 0, 3}, []int{1, 12, 0, 3}},
		{2, 2, []int{1, 12, 0, 3}, []int{1, 12, 2, 3}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Replacing %d at position %d", tt.value, tt.postition)
		t.Run(testname, func(t *testing.T) {
			ans := replaceValueInArray(tt.value, tt.postition, tt.inputArray)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
