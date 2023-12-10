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

	var board Board
	board.elements = make(BoardElements)
	board.minX = 0
	board.minY = 0
	var start_pos Coord
	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		board.maxX = len(buff)
		for i, c := range buff {
			if c == '.' {
				continue
			}
			coord := Coord{X: i, Y: j}
			board.elements[coord] = c
			if c == 'S' {
				start_pos = coord
			}
		}
		j++
		board.maxY = j
	}
	board.drawMap()
	fmt.Println(start_pos)

	var actual_pos, prev_pos [2]Coord
	prev_pos = [2]Coord{start_pos, start_pos}
	actual_pos = board.getInitialPipePair(start_pos)

	count := 0
	for {
		count += 1
		if actual_pos[0] == actual_pos[1] {
			break
		}

		new_actual_pos := [2]Coord{board.getNextStep(actual_pos[0], prev_pos[0]), board.getNextStep(actual_pos[1], prev_pos[1])}

		prev_pos = actual_pos
		actual_pos = new_actual_pos
	}
	fmt.Println(count)
}

func (board *Board) getInitialPipePair(initial_pos Coord) [2]Coord {
	ret := make([]Coord, 0)

	var coord Coord
	var c rune
	var ok bool

	coord = Coord{X: initial_pos.X, Y: initial_pos.Y - 1}
	c, ok = board.elements[coord]
	if ok && (c == '|' || c == '7' || c == 'F') {
		ret = append(ret, coord)
	}

	coord = Coord{X: initial_pos.X, Y: initial_pos.Y + 1}
	c, ok = board.elements[coord]
	if ok && (c == '|' || c == 'J' || c == 'L') {
		ret = append(ret, coord)
	}

	coord = Coord{X: initial_pos.X + 1, Y: initial_pos.Y}
	c, ok = board.elements[coord]
	if ok && (c == '-' || c == '7' || c == 'J') {
		ret = append(ret, coord)
	}

	coord = Coord{X: initial_pos.X - 1, Y: initial_pos.Y}
	c, ok = board.elements[coord]
	if ok && (c == '-' || c == 'L' || c == 'F') {
		ret = append(ret, coord)
	}

	if len(ret) != 2 {
		panic("Invalid initial pipe pair")
	}

	return [2]Coord{ret[0], ret[1]}
}

func (board *Board) getNextStep(actual_pos, prev_pos Coord) Coord {
	var new_pos Coord
	actual_pipe := board.elements[actual_pos]
	switch actual_pipe {
	case '|':
		if actual_pos.Y-1 == prev_pos.Y {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y + 1}
		} else {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y - 1}
		}
	case '-':
		if actual_pos.X-1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X + 1, Y: actual_pos.Y}
		} else {
			new_pos = Coord{X: actual_pos.X - 1, Y: actual_pos.Y}
		}
	case 'L':
		if actual_pos.X+1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y - 1}
		} else {
			new_pos = Coord{X: actual_pos.X + 1, Y: actual_pos.Y}
		}
	case 'J':
		if actual_pos.X-1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y - 1}
		} else {
			new_pos = Coord{X: actual_pos.X - 1, Y: actual_pos.Y}
		}
	case '7':
		if actual_pos.X-1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y + 1}
		} else {
			new_pos = Coord{X: actual_pos.X - 1, Y: actual_pos.Y}
		}
	case 'F':
		if actual_pos.X+1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y + 1}
		} else {
			new_pos = Coord{X: actual_pos.X + 1, Y: actual_pos.Y}
		}
	default:
		panic("Unknown pipe")
	}
	return new_pos
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
				fmt.Print(" ")
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
