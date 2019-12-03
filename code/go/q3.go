package main

import "fmt"

//Pair struct
type Pair struct {
	a, b interface{}
}

const size = 20

func goUp(array [][]string, num int, currentCoords Pair) Pair {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		x--
		if array[x][y] != "." {
			array[x][y] = "X"
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "|"
		}
	}
	currentCoords = Pair{x, y}
	return currentCoords
}

func goDown(array [][]string, num int, currentCoords Pair) Pair {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		x++
		if array[x][y] != "." {
			array[x][y] = "X"
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "|"
		}
	}
	currentCoords = Pair{x, y}
	return currentCoords
}

func goRight(array [][]string, num int, currentCoords Pair) Pair {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		y++
		if array[x][y] != "." {
			array[x][y] = "X"
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "-"
		}
	}
	currentCoords = Pair{x, y}
	return currentCoords
}

func goLeft(array [][]string, num int, currentCoords Pair) Pair {
	x := currentCoords.a.(int)
	y := currentCoords.b.(int)
	for i := 0; i < num; i++ {
		y--
		if array[x][y] != "." {
			array[x][y] = "X"
		} else if i+1 == num {
			array[x][y] = "+"
		} else {
			array[x][y] = "-"
		}
	}
	currentCoords = Pair{x, y}
	return currentCoords
}

func printGrid(array [][]string) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func main() {
	// Initializing the shit
	gridSlice := make([][]string, size)
	crossSections := make([]Pair, 30)
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

	current := starting

	current = goRight(gridSlice, 8, current)
	current = goUp(gridSlice, 5, current)
	current = goLeft(gridSlice, 5, current)
	current = goDown(gridSlice, 3, current)

	current = starting

	current = goUp(gridSlice, 7, current)
	current = goRight(gridSlice, 6, current)
	current = goDown(gridSlice, 4, current)
	current = goLeft(gridSlice, 4, current)

	printGrid(gridSlice)

}
