package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-04-input.txt"
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	fmt.Println(fileContent)
	stringSplit := strings.Split(fileContent, delimeter)

	// To deal with the \n on the end of the file read in
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

func main() {
	println("Day 04")
	lines := ReadFile("\n")
	result := solvePassword(lines[0], 1)

	fmt.Printf("Day 04 part 1 result of how many passwords => %d \n", result)

	resultPart2 := solvePassword(lines[0], 2)

	fmt.Printf("Day 04 part 2 result of how many passwords => %d \n", resultPart2)
}

func solvePassword(content string, version int) int {
	// we receive a line at a time so we split it
	parts := strings.Split(content, "-")
	partsInts := convertArrayToInts(parts)

	// get start and end and convert to ints
	start := partsInts[0]
	end := partsInts[1]

	validPasswordsCount := 0

	for password := start; password < end; password++ {
		// fmt.Println(password)
		if version == 1 {
			if validatePassword(password) {
				validPasswordsCount++
			}
		} else {
			if validatePasswordPart2(password) {
				validPasswordsCount++
			}
		}
	}

	return validPasswordsCount
}

func validatePassword(password int) bool {
	// convert back to string to get len etc.
	passString := strconv.FormatInt(int64(password), 10)

	// must be 6 digits long
	if len(passString) != 6 {
		return false
	}

	// digits from left to right never decrease
	// do this naively via pairs
	if passString[0] > passString[1] ||
		passString[1] > passString[2] ||
		passString[2] > passString[3] ||
		passString[3] > passString[4] ||
		passString[4] > passString[5] {
		return false
	}

	// check that at least 2 digits are the same
	duplicates := make(map[rune]int, 6)

	for _, passCharacter := range passString {
		// fmt.Printf("passstring %s character %c starts at byte position %d\n", passString, passCharacter, index)
		duplicates[passCharacter]++
		if duplicates[passCharacter] > 1 {
			fmt.Println(passString)
			return true
		}
	}

	return false
}

func validatePasswordPart2(password int) bool {
	// convert back to string to get len etc.
	passString := strconv.FormatInt(int64(password), 10)

	// must be 6 digits long
	if len(passString) != 6 {
		return false
	}

	// digits from left to right never decrease
	// do this naively via pairs
	if passString[0] > passString[1] ||
		passString[1] > passString[2] ||
		passString[2] > passString[3] ||
		passString[3] > passString[4] ||
		passString[4] > passString[5] {
		return false
	}

	// check for duplicates
	duplicates := make(map[rune]int, 6)

	for _, passCharacter := range passString {
		duplicates[passCharacter]++
	}

	// check that at least one number is repeated
	for _, duplicate := range duplicates {
		// fmt.Printf("input=> %s  duplicates %q and value of duplicate ->  %d \n", passString, duplicates, duplicate)
		if duplicate == 2 {
			return true
		}
	}

	return false
}
