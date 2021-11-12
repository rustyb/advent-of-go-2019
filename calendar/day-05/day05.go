package main

import (
	"fmt"

	"github.com/rustyb/adventofcode2019/calendar/utils"
)

type Op struct {
	Op    int
	Mode1 bool
	Mode2 bool
	Mode3 bool
}

func LoadOp(opcode int) Op {
	mode3 := opcode/10000 > 0
	mode2 := opcode%10000/1000 > 0
	mode1 := opcode%1000/100 > 0

	// println(opcode, "mode3:", opcode/10000, "mode2:", opcode%10000/1000, "mode1:", opcode%1000/100)
	// // mode3 := (opcode/1000)%10 > 0
	// // mode2 := (opcode/1000)%10 > 0
	// // mode1 := (opcode/100)%10 > 0
	// parameters := fmt.Sprintf("%.2d", opcode/100)
	// mode3 := parameters[1:] == "0"
	// mode2 := opcode%10000/1000 > 0
	// mode1 := parameters[:1] == "0"
	op := opcode % 100
	return Op{op, mode1, mode2, mode3}
}

func LoadArg(value int, mode bool, r []int) int {
	fmt.Println("LoadArg = ", value, mode)
	if mode {
		return value
	}

	return r[value]
}

func (op *Op) Load1(ip int, r []int) int {
	return LoadArg(r[ip+1], op.Mode1, r)
}

func (op *Op) Load2(ip int, r []int) int {
	return LoadArg(r[ip+2], op.Mode2, r)
}

func main() {
	lines := utils.ReadFileLines(",", "day-05-input.txt")
	// fmt.Println(lines)
	inputNumbers := utils.ConvertArrayToInts(lines)
	// fmt.Println(inputNumbers)
	// result := runComputer(inputNumbers, []int{1})
	// fmt.Println(result)

	r2 := make([]int, len(inputNumbers))
	copy(r2, inputNumbers)

	resultP2 := runComputer(r2, []int{5})
	fmt.Println(resultP2)
}

// Run computer from day 02
func runComputer(inputArray []int, input []int) []int {
	// lets not mutate the original array but return a new one
	arrayAfterRun := []int{}
	// copy(arrayAfterRun, inputArray)

	indexPosition := 0

	for {
		op := LoadOp(inputArray[indexPosition])
		fmt.Println("ITS AND OP", op)
		switch op.Op {
		case 1, 2:
			fmt.Println("Current op code", op.Op, "indexPositions", indexPosition, "current val = ", inputArray[indexPosition])
			a := op.Load1(indexPosition, inputArray)
			b := op.Load2(indexPosition, inputArray)
			c := inputArray[indexPosition+3]
			if op.Op == 1 {
				inputArray[c] = a + b
			} else {
				inputArray[c] = a * b
			}
			indexPosition += 4
			fmt.Println("After Current op code", op.Op, "indexPositions", indexPosition, "current val = ", inputArray[a], input)

		case 3:
			fmt.Println("Current op code", op.Op, "indexPositions", indexPosition, "current val = ", inputArray[indexPosition])
			a := inputArray[indexPosition+1]
			inputArray[a], input = input[0], input[1:]
			indexPosition += 2
			fmt.Println("After Current op code", op.Op, "indexPositions", indexPosition, "current val = ", inputArray[a], input)

		case 4:
			a := op.Load1(indexPosition, inputArray)
			arrayAfterRun = append(arrayAfterRun, a)
			indexPosition += 2

		case 5:
			fmt.Println("Current op code", "indexPositions", indexPosition, op.Op, "current val = ", inputArray[indexPosition])
			a := op.Load1(indexPosition, inputArray)
			b := op.Load2(indexPosition, inputArray)
			println("5OPA -> ", a)
			println("5OPB -> ", b)
			indexPosition += 3
			if a != 0 {
				// indexPosition += 3
				indexPosition = b
			}

		case 6:
			fmt.Println("Current op code", op.Op, "indexPositions", indexPosition, "current val = ", inputArray[indexPosition], inputArray)
			a := op.Load1(indexPosition, inputArray)
			b := op.Load2(indexPosition, inputArray)
			println("6OPA -> ", a)
			println("6OPB -> ", b)
			indexPosition += 3
			if a == 0 {
				indexPosition = b
			}
			fmt.Println("After Current op code", op.Op, "indexPositions", indexPosition)

		case 7:
			fmt.Println("Current op code", op.Op, "indexPositions", indexPosition, "current val = ", inputArray[indexPosition])
			// a := op.Load1(indexPosition, inputArray)
			// b := op.Load2(indexPosition, inputArray)
			c := indexPosition + 3
			a := inputArray[indexPosition+1]
			b := inputArray[indexPosition+2]

			if a < b {
				inputArray[c] = 1
			} else {
				inputArray[c] = 0
			}
			indexPosition += 4
			fmt.Println("After Current op code", op.Op, "indexPositions", indexPosition)

		case 8:
			fmt.Println("Current op code", op.Op, "indexPositions", indexPosition, "current val = ", inputArray[indexPosition])

			// a := op.Load1(indexPosition, inputArray)
			// b := op.Load2(indexPosition, inputArray)
			// c := inputArray[indexPosition+3]
			a := inputArray[indexPosition+1]
			b := inputArray[indexPosition+2]
			fmt.Println("=====a", a)
			fmt.Println("=====b", b)
			// fmt.Println("=====c", c)
			if a == b {
				inputArray[indexPosition+3] = 1
			} else {
				inputArray[indexPosition+3] = 0
			}
			indexPosition += 4
			fmt.Println("After Current op code", op.Op, "indexPositions", indexPosition)

		case 99:
			return arrayAfterRun

		default:
			panic("invalid op code")
		}
		// var section []int

		// section = arrayAfterRun[i : i+4]
		// operation := section[0]

		// // // do addition
		// // if operation == 1 {
		// // 	input1 := section[1]
		// // 	input2 := section[2]
		// // 	resultPosition := section[3]
		// // 	arrayAfterRun[resultPosition] = arrayAfterRun[input1] + arrayAfterRun[input2]
		// // }

		// // // do multiplication
		// // if operation == 2 {
		// // 	input1 := section[1]
		// // 	input2 := section[2]
		// // 	resultPosition := section[3]
		// // 	arrayAfterRun[resultPosition] = arrayAfterRun[input1] * arrayAfterRun[input2]
		// // }

		// if operation == 99 {
		// 	return arrayAfterRun
		// }
	}
	return arrayAfterRun
}
