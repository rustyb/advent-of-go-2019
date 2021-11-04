package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	var tests = []struct {
		value int
		want  bool
	}{
		{1234567, false},
		{754632, false},
		{678889, true},
		{567891, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input %d", tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := validatePassword(tt.value)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestValidatePasswordPart2(t *testing.T) {
	var tests = []struct {
		value int
		want  bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Input %d", tt.value)
		t.Run(testname, func(t *testing.T) {
			ans := validatePasswordPart2(tt.value)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
