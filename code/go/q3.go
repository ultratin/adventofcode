package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

const (
	MaxUint uint = (1 << bits.UintSize) - 1
	MaxInt  int  = (1<<bits.UintSize)/2 - 1
	MinInt  int  = (1 << bits.UintSize) / -2
	size         = 100000
)

//Pair struct
type Pair struct {
	a, b interface{}
}

func goUp(array [][]string, num int, currentCoords *Pair, crossSections *[]Pair) {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		x--
		if array[x][y] != "." {
			array[x][y] = "X"
			*crossSections = append(*crossSections, Pair{x, y})
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "|"
		}
	}
	*currentCoords = Pair{x, y}
}

func goDown(array [][]string, num int, currentCoords *Pair, crossSections *[]Pair) {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		x++
		if array[x][y] != "." {
			array[x][y] = "X"
			*crossSections = append(*crossSections, Pair{x, y})
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "|"
		}
	}
	*currentCoords = Pair{x, y}
}

func goRight(array [][]string, num int, currentCoords *Pair, crossSections *[]Pair) {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		y++
		if array[x][y] != "." {
			array[x][y] = "X"
			*crossSections = append(*crossSections, Pair{x, y})
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "-"
		}
	}
	*currentCoords = Pair{x, y}
}

func goLeft(array [][]string, num int, currentCoords *Pair, crossSections *[]Pair) {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		y--
		if array[x][y] != "." {
			array[x][y] = "X"
			*crossSections = append(*crossSections, Pair{x, y})
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "-"
		}
	}
	*currentCoords = Pair{x, y}
}

func printGrid(array [][]string) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func calculateDist(grid [][]string, crossSections []Pair) int {
	smallest := MaxInt
	x1 := len(grid) - 2
	y1 := 1
	for i := range crossSections {
		x2 := crossSections[i].a.(int)
		y2 := crossSections[i].b.(int)
		distance := abs(x1-x2) + abs(y1-y2)
		smallest = IntMin(smallest, distance)
	}
	return smallest
}

// IntMin for min function for int
func IntMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

func main() {
	// Initializing
	gridSlice := make([][]string, size)
	crossSections := make([]Pair, 0)
	crossPtr := &crossSections
	for i := range gridSlice {
		gridSlice[i] = make([]string, size)
	}
	for i := range gridSlice {
		for j := range gridSlice[i] {
			gridSlice[i][j] = "."
		}
	}
	starting := Pair{len(gridSlice[0]) - 2, 1}
	gridSlice[starting.a.(int)][starting.b.(int)] = "O"

	file, err := os.Open("../input/q3.txt")
	if err != nil {
		log.Fatalf("open file error: &v", err)
		return
	}
	defer file.Close()
	current := &starting
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines := sc.Text()
		instructions := strings.Split(lines, ",")
		*current = Pair{len(gridSlice[0]) - 2, 1}
		for i := range instructions {
			direction := (instructions[i][0:1])
			steps, _ := strconv.Atoi((instructions[i])[1:])
			if direction == "U" {
				goUp(gridSlice, steps, current, crossPtr)
			} else if direction == "D" {
				goDown(gridSlice, steps, current, crossPtr)
			} else if direction == "R" {
				goRight(gridSlice, steps, current, crossPtr)
			} else {
				goLeft(gridSlice, steps, current, crossPtr)
			}
		}
	}

	fmt.Println(calculateDist(gridSlice, crossSections))
}
