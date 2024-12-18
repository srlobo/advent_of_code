package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	coords := make([]Coord, 0)
	boardMaxX := 71
	boardMaxY := 71
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		buffSplitted := strings.Split(buff, ",")
		i, _ := strconv.Atoi(buffSplitted[0])
		j, _ := strconv.Atoi(buffSplitted[1])

		coords = append(coords, Coord{X: i, Y: j})
	}

	initPoint := 0
	lastPoint := len(coords) - 1
	actualPoint := initPoint + ((lastPoint - initPoint) / 2)
	for {
		board := makeBoardFromCoordArray(coords[:actualPoint], boardMaxX, boardMaxY)
		_, ok := dijkstra(board, Coord{0, 0}, Coord{boardMaxX - 1, boardMaxY - 1})
		// board.drawMap()
		// fmt.Println("Testing", initPoint, lastPoint, "distance", d, ok)
		if ok {
			initPoint = actualPoint
		} else {
			lastPoint = actualPoint
		}
		actualPoint = initPoint + ((lastPoint - initPoint) / 2)

		if lastPoint-initPoint <= 1 {
			break
		}
	}
	fmt.Println("The point in the aray is: ", actualPoint, coords[actualPoint])
}

func makeBoardFromCoordArray(coord []Coord, maxX, maxY int) Board {
	board := Board{elements: make(BoardElements), minX: 0, minY: 0, maxX: maxX, maxY: maxY}
	for _, c := range coord {
		board.elements[c] = '#'
	}
	return board
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
