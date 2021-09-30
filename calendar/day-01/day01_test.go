package main

import (
	"fmt"
	"testing"
)

// Simple test to check that the calculateFuel method works
func TestCalculateFuel(t *testing.T) {
	ans := calculateFuel(12)
	if ans != 2 {
		t.Errorf("calculateFuel(12) = %d; want 2", ans)
	}
}

// Throw a bunch of test calculations at calculateFuel
func TestCalculateFuelTableDriven(t *testing.T) {
	var tests = []struct {
		mass, want int
	}{
		{12, 2},
		{81157, 27050},
		{80969, 26987},
		{113477, 37823},
		{81295, 27096},
		{70537, 23510},
		{90130, 30041},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Mass %d", tt.mass)
		t.Run(testname, func(t *testing.T) {
			ans := calculateFuel(tt.mass)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Test the calculateAdditionalFuel method
// should return 2 also

func TestCalculateAdditionalFuel(t *testing.T) {
	ans := calculateAdditionalFuel(12)
	if ans != 2 {
		t.Errorf("calculateAdditionalFuel(12) = %d; want 2", ans)
	}
}

func TestCalculateAdditionalFuelTableDriven(t *testing.T) {
	var tests = []struct {
		mass, want int
	}{
		{12, 2},
		{81157, 40546},
		{80969, 40453},
		{113477, 56703},
		{81295, 40617},
		{70537, 35236},
		{90130, 45032},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Mass %d", tt.mass)
		t.Run(testname, func(t *testing.T) {
			ans := calculateAdditionalFuel(tt.mass)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
