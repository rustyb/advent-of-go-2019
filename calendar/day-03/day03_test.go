package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAbsoluteValue(t *testing.T) {
	var tests = []struct {
		value int
		want  int
	}{
		{12, 12},
		{-12, 12},
		{0, 0},
		{-0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input %d", tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := Abs(tt.value)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestMinValue(t *testing.T) {
	var tests = []struct {
		value1, value2 int
		want           int
	}{
		{12, 3, 3},
		{-12, 12, -12},
		{0, 0, 0},
		{-0, 0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input %d %d", tt.value1, tt.value2)
		t.Run(testname, func(t *testing.T) {
			ans := Min(tt.value1, tt.value2)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	var tests = []struct {
		value XYPoint
		want  int
	}{
		{XYPoint{12, 3}, 15},
		{XYPoint{-12, 3}, 15},
		{XYPoint{-2, 3}, 5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input %v", tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := Distance(tt.value)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
