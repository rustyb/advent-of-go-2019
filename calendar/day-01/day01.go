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

func getFuelForEach(inputArray []int) []int {
	intArrayToReturn := []int{}

	for _, ele := range inputArray {
		moduleFuel := (ele / 3) - 2

		intArrayToReturn = append(intArrayToReturn, moduleFuel)
	}

	return intArrayToReturn
}

// Problem requires to divide every number by 3 then subtract 2 and round down.

func main() {
	fmt.Println("Hello, World!")
	stringArray := ReadFile("\n")
	intArray := convertArrayToInts(stringArray)
	fuelForEach := getFuelForEach(intArray)
	fmt.Printf("%d\n", fuelForEach)

	totalFuel := 0
	for _, num := range fuelForEach {
		totalFuel += num
	}

	fmt.Printf("Total Fuel required to launch %d\n", totalFuel)
}
