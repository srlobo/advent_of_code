package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
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

	buttonRegexp := regexp.MustCompile(`[^:]+: X.(\d+), Y.(\d+)`)

	tokens := float64(0)
	for fileScanner.Scan() {
		var buttonA, buttonB, buttonC Button
		buff := fileScanner.Text()
		res := buttonRegexp.FindStringSubmatch(buff)
		buttonA.X, _ = strconv.ParseFloat(res[1], 64)
		buttonA.Y, _ = strconv.ParseFloat(res[2], 64)

		fileScanner.Scan()
		buff = fileScanner.Text()
		res = buttonRegexp.FindStringSubmatch(buff)
		buttonB.X, _ = strconv.ParseFloat(res[1], 64)
		buttonB.Y, _ = strconv.ParseFloat(res[2], 64)

		fileScanner.Scan()
		buff = fileScanner.Text()
		res = buttonRegexp.FindStringSubmatch(buff)
		buttonC.X, _ = strconv.ParseFloat(res[1], 64)
		buttonC.Y, _ = strconv.ParseFloat(res[2], 64)

		fileScanner.Scan()
		b := ((buttonC.X * buttonA.Y) - (buttonA.X * buttonC.Y)) / ((buttonB.X * buttonA.Y) - (buttonB.Y * buttonA.X))
		a := (buttonC.Y - (b * buttonB.Y)) / buttonA.Y

		if a > 100 || b > 100 {
			continue
		}

		if float64(int(a)) != a || float64(int(b)) != b {
			continue
		}

		fmt.Printf("Button A: %v\n", buttonA)
		fmt.Printf("Button B: %v\n", buttonB)
		fmt.Printf("Button C: %v\n", buttonC)
		fmt.Printf("a: %v\n", a)
		fmt.Printf("b: %v\n", b)
		fmt.Println("--------------------")

		tokens += 3*a + b

	}
	fmt.Printf("Tokens: %v\n", tokens)
}

type Button struct {
	X float64
	Y float64
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dijkstra(board Board, start, end Coord) int {
	getNeighbours := func(coord Coord) map[Coord]int {
		ret := make(map[Coord]int)
		var c Coord

		// left
		c = Coord{X: coord.X - 1, Y: coord.Y}
		ret[c], _ = strconv.Atoi(string(board.elements[c]))

		// right
		c = Coord{X: coord.X + 1, Y: coord.Y}
		ret[c], _ = strconv.Atoi(string(board.elements[c]))

		// up
		c = Coord{X: coord.X, Y: coord.Y - 1}
		ret[c], _ = strconv.Atoi(string(board.elements[c]))

		// down
		c = Coord{X: coord.X, Y: coord.Y + 1}
		ret[c], _ = strconv.Atoi(string(board.elements[c]))

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
		currentNode = nextNode
	}
	return visitedSet[end]
}
