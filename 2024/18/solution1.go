package main

import (
	"bufio"
	"fmt"
	"math"
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

	board := Board{elements: make(BoardElements), minX: 0, minY: 0, maxX: 71, maxY: 71}

	count := 0
	for fileScanner.Scan() {
		if count == 1024 {
			break
		}
		buff := fileScanner.Text()
		buffSplitted := strings.Split(buff, ",")
		i, _ := strconv.Atoi(buffSplitted[0])
		j, _ := strconv.Atoi(buffSplitted[1])
		board.elements[Coord{X: i, Y: j}] = '#'
		count++
	}
	board.drawMap()

	fmt.Println(dijkstra(board, Coord{X: 0, Y: 0}, Coord{X: 70, Y: 70}))
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

func dijkstra(board Board, start, end Coord) int {
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
		min := math.MaxInt32
		var ret Coord
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
			unvisitedSet[Coord{X: i, Y: j}] = math.MaxInt32
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
			break
		}
		nextNode := getNextNode(unvisitedSet)
		// fmt.Println("Next node: ", nextNode)
		// fmt.Println("Visited set: ", visitedSet)
		// fmt.Println("Unvisited set: ", unvisitedSet)
		currentNode = nextNode
	}
	return visitedSet[end]
}
