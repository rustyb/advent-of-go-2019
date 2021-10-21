package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func ReadFile(delimeter string) []string {
	filePath := "day-03-input.txt"
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

// lets define our data structures
// we have the most basic an xy point
// a wire is an array of these xy points

type XYPoint struct {
	X, Y int
}

type Wire struct {
	data map[XYPoint]int
}

func (w Wire) Get(x, y int) int {
	v, ok := w.data[XYPoint{x, y}]
	if ok {
		return v
	}
	return 0
}

// method to set wiresections in a Wire
func (w *Wire) Set(x, y, d int) {
	xy := XYPoint{x, y}
	// println("Adding" + strconv.Itoa(xy.X) + "," + strconv.Itoa(xy.Y))

	v, ok := w.data[xy]
	if ok {
		d = Min(d, v)
	}
	w.data[xy] = d
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Distance(p XYPoint) int {
	return Abs(p.X) + Abs(p.Y)
}

// find the crossing point
func Crossings(w1, w2 Wire) []XYPoint {
	result := []XYPoint{}
	for k := range w1.data {
		if _, ok := w2.data[k]; ok {
			if k.X == 0 && k.Y == 0 {
				continue
			}
			println(k.X, k.Y)
			result = append(result, k)
		}
	}
	println(result)
	return result
}

func LoadWire(txt string) Wire {
	w := Wire{}
	w.data = make(map[XYPoint]int)

	directions := strings.Split(txt, ",")

	// starting point init at XY 0,0 and set distance 0
	x, y, d := 0, 0, 0
	// fmt.Printf("start point %d,%d and distance %d ", x, y, d)

	// function to add wire sections to this wire route
	addWireSection := func(vector *int, directionVector int, amount int) {
		stop := *vector + directionVector*(amount+1)
		for *vector += directionVector; *vector != stop; *vector += directionVector {
			d++
			w.Set(x, y, d)
		}
		*vector -= directionVector
	}

	// iterate over the directions to construct the wire

	for _, direction := range directions {
		// way to turn is first character
		goDirection := direction[0]

		// convert the distance to an int
		amount, err := strconv.Atoi(direction[1:])

		if err != nil {
			panic(err)
		}

		// fmt.Printf("%v We are going %v by %d met.\n", direction, goDirection, amount)
		switch goDirection {
		case 'R':
			// fmt.Printf("Moving right\n")
			addWireSection(&x, 1, amount)
		case 'L':
			// fmt.Printf("Moving left\n")
			addWireSection(&x, -1, amount)
		case 'U':
			// fmt.Printf("Moving up\n")
			addWireSection(&y, 1, amount)
		case 'D':
			// fmt.Printf("Moving down\n")
			addWireSection(&y, -1, amount)
		}
	}

	return w
}

func CrossingsShortestDistance(w1, w2 Wire) int {
	cs := Crossings(w1, w2)
	d := Distance(cs[0])
	for _, c := range cs {
		d = Min(d, Distance(c))
	}
	return d
	// println("result of crossings", cs)

	// return 1
}

func CrossingsShortestSignal(w1, w2 Wire) int {
	result := math.MaxInt32
	for k := range w1.data {
		if v2, ok := w2.data[k]; ok {
			if k.X == 0 && k.Y == 0 {
				continue
			}
			result = Min(result, w1.data[k]+v2)
		}
	}
	return result
}

func main() {
	println("Day 03")
	wires := ReadFile("\n")
	println(wires)

	w1 := LoadWire(wires[0])
	w2 := LoadWire(wires[1])

	shortestD := CrossingsShortestDistance(w1, w2)
	quickestD := CrossingsShortestSignal(w1, w2)

	println("shortest =", shortestD)
	println("shortest =", quickestD)

}
