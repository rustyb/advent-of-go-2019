package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-02-input.txt"
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

// Once you have a working computer, the first step is to restore the
// gravity assist program (your puzzle input) to the "1202 program alarm"
// state it had just before the last computer caught fire. To do this,
// before running the program, replace position 1 with the value 12 and
// replace position 2 with the value 2. What value is left at position 0 after the program halts?
func replaceValueInArray(value int, postition int, inputArray []int) []int {
	newArray := make([]int, len(inputArray))
	copy(newArray, inputArray)
	newArray[postition] = value
	return newArray
}

func runComputer(inputArray []int) []int {
	// lets not mutate the original array but return a new one
	arrayAfterRun := make([]int, len(inputArray))
	copy(arrayAfterRun, inputArray)

	for i := 0; i < len(arrayAfterRun); i += 4 {
		var section []int

		section = arrayAfterRun[i : i+4]
		operation := section[0]
		fmt.Println("INSTR", i, section, operation)

		// do addition
		if operation == 1 {
			input1 := section[1]
			input2 := section[2]
			resultPosition := section[3]
			arrayAfterRun[resultPosition] = arrayAfterRun[input1] + arrayAfterRun[input2]
			// fmt.Println("add", input1+input2, "to pos:", resultPosition, arrayAfterRun)
		}

		// do multiplication
		if operation == 2 {
			input1 := section[1]
			input2 := section[2]
			resultPosition := section[3]
			arrayAfterRun[resultPosition] = arrayAfterRun[input1] * arrayAfterRun[input2]
		}

		if operation == 99 {
			return arrayAfterRun
		}
	}
	return arrayAfterRun
}

func main() {
	strings := ReadFile(",")
	inputNumbers := convertArrayToInts(strings)
	fmt.Printf("Welcome to day 2")
	// fmt.Printf("%d\n", inputNumbers)

	// result := FixRun(12, 2, inputNumbers)
	// fmt.Println(result)
	// first part replace values in array
	fixedArrayStep1 := replaceValueInArray(12, 1, inputNumbers)
	fmt.Printf("%d\n", fixedArrayStep1)
	fixedArrayStep2 := replaceValueInArray(2, 2, fixedArrayStep1)
	fmt.Printf("%d\n", fixedArrayStep2)

	// calculate the result
	resultArray := runComputer(fixedArrayStep1)

	fmt.Printf("Final result = > %d\n", resultArray[0])
}
