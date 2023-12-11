package main

import (
	"bufio"
	"fmt"
	"os"
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

	board := Board{elements: make(BoardElements)}
	expansion_factor := 999999
	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()

		empty_line := true
		for i, c := range buff {
			coord := Coord{X: i, Y: j}
			if c == '#' {
				board.elements[coord] = c
				empty_line = false
			}
			board.maxX = i
		}
		if empty_line {
			j += expansion_factor
		}
		j++
		board.maxY = j - 1
	}
	board.minX = 0
	board.minX = 0
	fmt.Println(len(board.elements))
	for i := 0; i <= board.maxX; i++ {
		empty_col := true
		for j := 0; j <= board.maxY; j++ {
			coord := Coord{X: i, Y: j}
			if _, ok := board.elements[coord]; ok {
				empty_col = false
			}
		}
		if empty_col {
			elements_to_change := make([]Coord, 0)
			for k := range board.elements {
				elements_to_change = append(elements_to_change, k)
			}
			for _, k := range elements_to_change {
				if k.X > i {
					v := board.elements[k]
					delete(board.elements, k)
					// fmt.Println("Coord ", k, " will be changed to ", Coord{X: k.X + expansion_factor, Y: k.Y})
					k.X += expansion_factor
					board.elements[k] = v
				}
			}
			board.maxX += expansion_factor
			i += expansion_factor
		}
	}
	fmt.Println(len(board.elements))

	// board.drawMap()

	remaining_elements := board.elements
	total := 0
	for {
		if len(remaining_elements) == 0 {
			break
		}
		var current_element Coord
		for current_element = range remaining_elements { // Choose one random element from tiles_to_process
			break
		}
		delete(remaining_elements, current_element)
		for element := range remaining_elements {
			dist := calculateDistance(current_element, element)
			// fmt.Println("Distance between ", current_element, " and ", element, " is ", dist)
			total += dist
		}
	}

	fmt.Println(total)
}

func calculateDistance(a, b Coord) int {
	dist := abs(b.X-a.X) + abs(b.Y-a.Y)
	return dist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Coord struct {
	X int
	Y int
}

type BoardElements map[Coord]rune

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
				fmt.Print(string(dot))
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
