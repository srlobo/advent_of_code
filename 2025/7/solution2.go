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

	board := Board{minX: 0, minY: 0, elements: make(BoardElements)}
	var start Coord

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	y := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		x := 0
		for _, c := range buff {
			if c == 'S' {
				start = Coord{x, y}
				board.elements[start] = c
			} else if c == '^' {
				coord := Coord{x, y}
				board.elements[coord] = c
			}
			x++
			board.maxX = x
		}
		y++
	}
	board.maxY = y

	board.drawMap()
	// Beams hold the x coord of the beam, and how many beams lead to this point
	beams := make(map[int]int)
	beams[start.X] = 1
	for row := start.Y; row < board.maxY; row++ {
		currentBeams := make(map[int]int)
		for beam, n := range beams {
			coord := Coord{beam, row}
			if board.elements[coord] == '^' {
				currentBeams[beam+1] += n
				currentBeams[beam-1] += n
			} else {
				currentBeams[beam] += n
			}
		}
		beams = currentBeams
	}
	total := 0
	for _, n := range beams {
		total += n
	}
	fmt.Println(total)
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

func (board *Board) drawMapWithCoords(coords map[Coord]byte) {
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
			beam, ok2 := coords[c]
			if ok && ok2 {
				fmt.Print(string(beam))
			} else if ok && !ok2 {
				fmt.Print(string(dot))
			} else if !ok && ok2 {
				fmt.Print(string(beam))
			} else if !ok && !ok2 {
				fmt.Print(".")
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

func dijkstra(board Board, start, end Coord) (int, bool) {
	getNeighbours := func(coord Coord) map[Coord]int {
		ret := make(map[Coord]int)
		var c Coord

		// left
		c = Coord{X: coord.X - 1, Y: coord.Y}
		if c.X >= board.minX && c.X < board.maxX && c.Y >= board.minY && c.Y < board.maxY {
			if _, ok := board.elements[c]; !ok {
				ret[c] = 1
			}
		}

		// right
		c = Coord{X: coord.X + 1, Y: coord.Y}
		if c.X >= board.minX && c.X < board.maxX && c.Y >= board.minY && c.Y < board.maxY {
			if _, ok := board.elements[c]; !ok {
				ret[c] = 1
			}
		}

		// up
		c = Coord{X: coord.X, Y: coord.Y - 1}
		if c.X >= board.minX && c.X < board.maxX && c.Y >= board.minY && c.Y < board.maxY {
			if _, ok := board.elements[c]; !ok {
				ret[c] = 1
			}
		}

		// down
		c = Coord{X: coord.X, Y: coord.Y + 1}
		if c.X >= board.minX && c.X < board.maxX && c.Y >= board.minY && c.Y < board.maxY {
			if _, ok := board.elements[c]; !ok {
				ret[c] = 1
			}
		}

		return ret
	}

	getNextNode := func(unvisitedSet map[Coord]int) Coord {
		min := MaxInt
		var ret Coord
		for coord := range unvisitedSet { // Init the ret variable
			ret = coord
			break
		}
		for coord, v := range unvisitedSet {
			if v < min {
				min = v
				ret = coord
			}
		}
		return ret
	}

	unvisitedSet := make(map[Coord]int)
	visitedSet := make(map[Coord]int)
	for j := 0; j < board.maxY; j++ {
		for i := 0; i < board.maxX; i++ {
			unvisitedSet[Coord{X: i, Y: j}] = MaxInt
		}
	}
	unvisitedSet[start] = 0
	currentNode := start
	for {
		currentNodeDistance := unvisitedSet[currentNode]
		possibleNeighbourgs := getNeighbours(currentNode)
		for node, distance := range possibleNeighbourgs {
			if _, ok := visitedSet[node]; ok {
				continue
			}
			if unvisitedSet[node] > currentNodeDistance+distance {
				unvisitedSet[node] = currentNodeDistance + distance
			}
		}
		delete(unvisitedSet, currentNode)
		visitedSet[currentNode] = currentNodeDistance
		if _, ok := visitedSet[end]; ok {
			return visitedSet[end], true
		}
		if len(unvisitedSet) == 0 {
			return visitedSet[currentNode], false
		}
		nextNode := getNextNode(unvisitedSet)
		if unvisitedSet[nextNode] == MaxInt {
			return visitedSet[currentNode], false
		}

		currentNode = nextNode
	}
}
