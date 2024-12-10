package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var empty = struct{}{}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		defer readFile.Close()
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	j := 0
	board := Board{elements: make(BoardElements), minX: 0, minY: 0}

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		board.maxX = len(buff) - 1
		for i := 0; i < len(buff); i++ {
			if r := buff[i]; r != '.' {
				board.elements[Coord{X: i, Y: j}], _ = strconv.Atoi(string(r))
			}
		}
		j++
	}
	board.maxY = j - 1
	board.drawMap()

	total := 0
	for j := 0; j <= board.maxY; j++ {
		for i := 0; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			if height, ok := board.elements[c]; ok && height == 0 {
				rating := board.findTrail(c)
				fmt.Println("trailHead:", c, rating)
				total += rating
			}
		}
	}
	fmt.Println(total)
}

func (board *Board) findTrail(c Coord) int {
	// Calculamos cada direccion de trail (u d l r) y llamamos a esta función de forma recursiva hasta encontrar el 0
	var newCoord Coord
	var newValue int
	total := 0
	actualValue := board.elements[c]
	if actualValue == 9 {
		return 1
	}

	// up
	newCoord = Coord{X: c.X, Y: c.Y - 1}

	newValue = board.elements[newCoord]
	if newValue-actualValue == 1 {
		total += board.findTrail(newCoord)
	}

	// down
	newCoord = Coord{X: c.X, Y: c.Y + 1}

	newValue = board.elements[newCoord]
	if newValue-actualValue == 1 {
		total += board.findTrail(newCoord)
	}

	// left
	newCoord = Coord{X: c.X - 1, Y: c.Y}

	newValue = board.elements[newCoord]
	if newValue-actualValue == 1 {
		total += board.findTrail(newCoord)
	}

	// right
	newCoord = Coord{X: c.X + 1, Y: c.Y}

	newValue = board.elements[newCoord]
	if newValue-actualValue == 1 {
		total += board.findTrail(newCoord)
	}

	return total
}

type Coord struct {
	X int
	Y int
}

type (
	BoardElements map[Coord]int
	trailTails    map[Coord]struct{}
)

type Board struct {
	elements BoardElements

	minX int
	minY int
	maxX int
	maxY int
}

func (board *Board) drawMap() {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)

	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			if ok {
				fmt.Print(dot)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (board *Board) drawMapWithCoords(coords map[Coord]struct{}) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	minY := MaxInt
	maxY := 0

	for c := range coords {
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	for j := minY; j <= maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			_, ok2 := coords[c]
			if ok && ok2 {
				fmt.Print(string(dot))
			} else if ok && !ok2 {
				fmt.Print(string(dot))
			} else if !ok && ok2 {
				fmt.Print(string("."))
			} else if !ok && !ok2 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
