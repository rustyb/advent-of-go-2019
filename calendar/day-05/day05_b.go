package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	slice := strings.Split(sc.Text(), ",")

	var instructions []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		instructions = append(instructions, n)
	}

	instructions = append(instructions, 0, 0)
	var programCounter int

	for programCounter > -1 {
		opCode := instructions[programCounter] % 100
		parameters := fmt.Sprintf("%.2d", instructions[programCounter]/100)

		a := instructions[programCounter+1]
		b := instructions[programCounter+2]
		if opCode != 3 && opCode != 4 && opCode != 99 {
			if parameters[1:] == "0" {
				a = instructions[a]
			}
			if parameters[:1] == "0" {
				b = instructions[b]
			}
		}

		switch opCode {
		case 1:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
			instructions[instructions[programCounter+3]] = a + b
			programCounter += 4 //Number of instructions
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 2:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
			instructions[instructions[programCounter+3]] = a * b
			programCounter += 4 //Number of instructions
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 3:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
			sc.Scan()
			instructions[instructions[programCounter+1]], _ = strconv.Atoi(sc.Text())
			programCounter += 2 //Number of instructions
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 4:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
			fmt.Println(instructions[a])
			programCounter += 2 //Number of instructions
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 5:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
			programCounter += 3
			if a != 0 {
				programCounter = b //Jump
			}
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 6:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])

			programCounter += 3
			if a == 0 {
				programCounter = b //Jump
			}
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 7:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
			if a < b {
				instructions[instructions[programCounter+3]] = 1
			} else {
				instructions[instructions[programCounter+3]] = 0
			}
			programCounter += 4 //Number of instructions
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 8:
			fmt.Println("Current op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
			fmt.Println("=====a", a)
			fmt.Println("=====b", b)
			fmt.Println("=====instructions[programCounter+3]", instructions[programCounter+3])

			if a == b {
				instructions[instructions[programCounter+3]] = 1
			} else {
				instructions[instructions[programCounter+3]] = 0
			}
			programCounter += 4 //Number of instructions
			fmt.Println("===After op code", opCode, "indexPositions", programCounter, "current val = ", instructions[programCounter])
		case 99:
			programCounter = -1 //Exit
		default:
			fmt.Println("Error at instruction number ", programCounter)
		}
	}
}
