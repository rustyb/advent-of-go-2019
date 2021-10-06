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

// Test that setting the noun and verb actually sets the array
func TestSetNounVerbInArray(t *testing.T) {
	var tests = []struct {
		noun, verb int
		inputArray []int
		want       []int
	}{
		{12, 2, []int{1, 0, 0, 3}, []int{1, 12, 2, 3}},
		{2, 2, []int{1, 12, 0, 3}, []int{1, 2, 2, 3}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Using noun %d, verb %d with array %d", tt.noun, tt.verb, tt.inputArray)
		t.Run(testname, func(t *testing.T) {
			ans := setNounVerbInArray(tt.noun, tt.verb, tt.inputArray)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, wanted %d", ans, tt.want)
			}
		})
	}
}

// Test the run computer function
func TestRunComputer(t *testing.T) {
	var tests = []struct {
		inputArray []int
		want       []int
	}{
		{[]int{1, 0, 0, 3}, []int{1, 0, 0, 2}},
		// need to add more complex tests for this one
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Running the computer with %d", tt.inputArray)
		t.Run(testname, func(t *testing.T) {
			ans := runComputer(tt.inputArray)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
