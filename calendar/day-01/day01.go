package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-01-input.txt"
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	fmt.Println(fileContent)
	stringSplit := strings.Split(fileContent, delimeter)

	// To deal with the \n on the end of the file read in
	fmt.Printf("%q\n", stringSplit[:len(stringSplit)-1])
	return stringSplit[:len(stringSplit)-1]
}

// convert the string array to ints to be able to
func convertArrayToInts(inputArray []string) []int {
	intArrayToReturn := []int{}

	for _, ele := range inputArray { // you can escape index by _ keyword
		convertedString, err := strconv.Atoi(ele)

		if err != nil {
			panic(err)
		}

		intArrayToReturn = append(intArrayToReturn, convertedString)
	}
	return intArrayToReturn
}

/**
  The first step required to calculate the fuel following the formula
  fuel = mass / 3 - 2
**/

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

/*
  The second step requires to calculate the amount of extra fuel required to
  carry the extra fuel. Treat  treat the fuel amount you just calculated as
  the input mass and repeat the process, continuing until a fuel requirement
  is zero or negative.
*/

func calculateAdditionalFuel(mass int) int {
	totalFuel := 0

	for {
		fuel := calculateFuel(mass)
		if fuel <= 0 {
			return totalFuel
		}
		// still more fuel to carry
		totalFuel += fuel

		mass = calculateFuel(fuel)
		if mass <= 0 {
			return totalFuel
		}
		totalFuel += mass
	}
}

// Problem requires to divide every number by 3 then subtract 2 and round down.

func main() {
	stringArray := ReadFile("\n")
	intArray := convertArrayToInts(stringArray)

	totalFuel := 0
	for _, num := range intArray {
		fuel := calculateFuel(num)
		totalFuel += fuel
	}

	// Part 2 - calculate additional fuel required to carry fuel weight

	sumFuelWithAdditional := 0

	for _, num := range intArray {
		fuel := calculateAdditionalFuel(num)
		sumFuelWithAdditional += fuel
	}

	fmt.Printf("Total Fuel required to launch %d\n", totalFuel)
	fmt.Printf("Total Fuel + Fuel for fuel weight required to launch %d\n", sumFuelWithAdditional)
}
