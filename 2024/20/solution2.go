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

	board := Board{elements: make(BoardElements), minX: 0, minY: 0}
	var start, end Coord
	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		board.maxX = len(buff)
		for i := 0; i < len(buff); i++ {
			if buff[i] == '#' {
				board.elements[Coord{X: i, Y: j}] = '#'
			} else if buff[i] == 'S' {
				start = Coord{X: i, Y: j}
			} else if buff[i] == 'E' {
				end = Coord{X: i, Y: j}
			}
		}
		j++
	}
	board.maxY = j
	res, path, _ := dijkstra(board, start, end)
	fmt.Println(res)
	fmt.Println(path)
	pathMap := make(map[Coord]struct{})
	for _, c := range path {
		pathMap[c] = empty
	}
	// board.drawMapWithCoords(pathMap)

	distances := make(map[int]int)

	count := 0
	for i := 0; i < len(path); i++ {
		for j := i + 2; j < len(path); j++ {
			d := distance(path[i], path[j])
			if d <= 20 && j-i-d >= 100 {
				count++
				fmt.Println(path[i], path[j], d, j-i-d)
				distances[j-i-d]++
			}

		}
	}
	for i := 0; i < len(path); i++ {
		if d, ok := distances[i]; ok {
			fmt.Println(i, ":", d)
			continue
		}
	}
	fmt.Println(count)
}

func distance(a, b Coord) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
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

func makeBoardFromCoordArray(coord []Coord, maxX, maxY int) Board {
	board := Board{elements: make(BoardElements), minX: 0, minY: 0, maxX: maxX, maxY: maxY}
	for _, c := range coord {
		board.elements[c] = '#'
	}
	return board
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
	minY := board.minY
	maxY := board.maxY

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
				fmt.Print(string("O"))
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

func dijkstra(board Board, start, end Coord) (int, []Coord, bool) {
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
	previousCoord := make(map[Coord]Coord)
	var success bool
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
			success = true
			break
		}
		if len(unvisitedSet) == 0 {
			success = false
			break
		}
		nextNode := getNextNode(unvisitedSet)
		if unvisitedSet[nextNode] == MaxInt {
			success = false
			break
		}

		previousCoord[nextNode] = currentNode
		currentNode = nextNode
	}

	path := make([]Coord, 0)
	if success {
		currentNode = end
		for currentNode != start {
			path = append(path, currentNode)
			currentNode = previousCoord[currentNode]
		}
		path = append(path, start)

		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}

		return visitedSet[end], path, true
	} else {
		return visitedSet[currentNode], path, false
	}
}
