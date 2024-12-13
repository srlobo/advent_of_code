package main

import (
	"bufio"
	"fmt"
	"math"
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

	board := Board{elements: make(BoardElements), minX: 0, minY: 0}

	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		board.maxX = len(buff)
		for i, r := range buff {
			c := Coord{X: i, Y: j}
			board.elements[c] = r
		}
		j++
	}
	board.maxY = j

	visited := make(map[Coord]struct{})
	discovered := make(map[Coord]struct{})
	total := 0
	for n, r := range board.elements {
		if _, ok := visited[n]; ok {
			continue
		}
		fmt.Println("Starting group search from:", string(r))
		discovered[n] = empty
		group := make(map[Coord]struct{})
		for {
			if len(discovered) == 0 {
				break
			}
			for c := range discovered {
				if _, ok := visited[c]; ok {
					continue
				}
				delete(discovered, c)
				visited[c] = empty
				group[c] = empty
				for discoveredNode := range expand(board, visited, c) {
					// fmt.Println("Adding discovered node:", discoveredNode)
					discovered[discoveredNode] = empty
				}
			}
		}
		// if r != 'J' {
		// 	continue
		// }
		sides := getSides(board, group)
		fmt.Println("Group found:", string(r), len(group), "sides:", sides)
		total += len(group) * sides
	}
	fmt.Println("Total:", total)
}

func getSides(board Board, group map[Coord]struct{}) int {
	ret := 0
	minX, maxX, minY, maxY := getBoundingBox(group)
	// fmt.Println("minY:", minY, "maxY:", maxY, "minX:", minX, "maxX:", maxX)
	j := minY - 1
	for {
		j++
		if j > maxY {
			break
		}
		i := minX - 1
		for {
			i++
			if i > maxX {
				break
			}
			coord := Coord{X: i, Y: j}
			// fmt.Println("Checking coord:", coord)
			if _, ok := group[coord]; !ok {
				// fmt.Println("Skipping coord:", coord)
				// fmt.Println(group)
				continue
			}
			// We're just treating a component of the group, and moving horizontally. Let's see how many non-contiguous regions of frontiers exist
			otherSideCoord := Coord{X: i, Y: j - 1}
			// fmt.Println("Initial check frontier coord:", coord, otherSideCoord)
			if board.elements[coord] != board.elements[otherSideCoord] { // it's a frontier
				for board.elements[coord] != board.elements[otherSideCoord] && i <= maxX {
					i++
					coord = Coord{X: i, Y: j}
					if _, ok := group[coord]; !ok {
						break
					}

					otherSideCoord = Coord{X: i, Y: j - 1}
					// fmt.Println("Checking frontier coord:", coord, otherSideCoord)
				}
				i--
				ret++
			}
		}
	}
	fmt.Println("Sides: u", ret)

	j = minY - 1
	for {
		j++
		if j > maxY {
			break
		}
		i := minX - 1
		for {
			i++
			if i > maxX {
				break
			}
			coord := Coord{X: i, Y: j}
			// fmt.Println("Checking coord:", coord)
			if _, ok := group[coord]; !ok {
				continue
			}
			// We're just treating a component of the group, and moving horizontally. Let's see how many non-contiguous regions of frontiers exist
			otherSideCoord := Coord{X: i, Y: j + 1}
			// fmt.Println("Initial check frontier coord:", coord, otherSideCoord)
			if board.elements[coord] != board.elements[otherSideCoord] { // it's a frontier
				for board.elements[coord] != board.elements[otherSideCoord] && i <= maxX {
					i++
					coord = Coord{X: i, Y: j}
					if _, ok := group[coord]; !ok {
						break
					}

					otherSideCoord = Coord{X: i, Y: j + 1}
					// fmt.Println("Checking frontier coord:", coord, otherSideCoord)
				}
				i--
				ret++
			}
		}
	}
	fmt.Println("Sides: d", ret)

	i := minX - 1
	for {
		i++
		if i > maxX {
			break
		}
		j := minY - 1
		for {
			j++
			if j > maxY {
				break
			}
			coord := Coord{X: i, Y: j}
			// fmt.Println("Checking coord:", coord)
			if _, ok := group[coord]; !ok {
				// fmt.Println("Skipping coord:", coord)
				// fmt.Println(group)
				continue
			}
			// We're just treating a component of the group, and moving horizontally. Let's see how many non-contiguous regions of frontiers exist
			otherSideCoord := Coord{X: i - 1, Y: j}
			// fmt.Println("Initial check frontier coord:", coord, otherSideCoord, board.elements[coord] != board.elements[otherSideCoord])
			if board.elements[coord] != board.elements[otherSideCoord] { // it's a frontier
				for board.elements[coord] != board.elements[otherSideCoord] && j <= maxY {
					j++
					coord = Coord{X: i, Y: j}
					otherSideCoord = Coord{X: i - 1, Y: j}
					// fmt.Println("Checking frontier coord:", coord, otherSideCoord, board.elements[coord] != board.elements[otherSideCoord], board.elements[coord] != board.elements[otherSideCoord] || j <= maxY)
					if _, ok := group[coord]; !ok {
						// fmt.Println("Out")
						break
					}

				}
				j--
				ret++
			}
		}
	}
	fmt.Println("Sides: l", ret)

	i = minX - 1
	for {
		i++
		if i > maxX {
			break
		}
		j := minY - 1
		for {
			j++
			if j > maxY {
				break
			}
			coord := Coord{X: i, Y: j}
			// fmt.Println("Checking coord:", coord)
			if _, ok := group[coord]; !ok {
				// fmt.Println("Skipping coord:", coord)
				// fmt.Println(group)
				continue
			}
			// We're just treating a component of the group, and moving horizontally. Let's see how many non-contiguous regions of frontiers exist
			otherSideCoord := Coord{X: i + 1, Y: j}
			// fmt.Println("Initial check frontier coord:", coord, otherSideCoord)
			if board.elements[coord] != board.elements[otherSideCoord] { // it's a frontier
				for board.elements[coord] != board.elements[otherSideCoord] && j <= maxY {
					j++
					coord = Coord{X: i, Y: j}
					if _, ok := group[coord]; !ok {
						break
					}

					otherSideCoord = Coord{X: i + 1, Y: j}
					// fmt.Println("Checking frontier coord:", coord, otherSideCoord)
				}
				j--
				ret++
			}
		}
	}
	fmt.Println("Sides: r", ret)

	return ret
}

func getBoundingBox(elements map[Coord]struct{}) (minX, maxX, minY, maxY int) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	minY = MaxInt
	minX = MaxInt
	maxX = 0
	maxY = 0

	for c := range elements {
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
	}

	return minX, maxX, minY, maxY
}

func expand(board Board, visited map[Coord]struct{}, start Coord) map[Coord]struct{} {
	c := board.elements[start]
	result := make(map[Coord]struct{})
	var ok bool

	var coord Coord

	coord = Coord{X: start.X - 1, Y: start.Y}
	_, ok = visited[coord]
	if board.elements[coord] == c && !ok {
		result[coord] = empty
	}
	coord = Coord{X: start.X + 1, Y: start.Y}
	_, ok = visited[coord]
	if board.elements[coord] == c && !ok {
		result[coord] = empty
	}
	coord = Coord{X: start.X, Y: start.Y + 1}
	_, ok = visited[coord]
	if board.elements[coord] == c && !ok {
		result[coord] = empty
	}
	coord = Coord{X: start.X, Y: start.Y - 1}
	_, ok = visited[coord]
	if board.elements[coord] == c && !ok {
		result[coord] = empty
	}

	return result
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
