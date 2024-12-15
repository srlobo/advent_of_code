package main

import (
	"bufio"
	"fmt"
	"os"
)

var empty = struct{}{}

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
	var instructions string
	j := 0
	var botPosition Coord

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		if buff == "" {
			break
		}
		board.maxX = 2 * len(buff)
		for i := 0; i < len(buff); i++ {
			if buff[i] != '.' {
				coord := Coord{X: i * 2, Y: j}
				if buff[i] == '@' {
					botPosition = coord
				} else if buff[i] == '#' {
					board.elements[coord] = rune(buff[i])
					coord = Coord{X: i*2 + 1, Y: j}
					board.elements[coord] = rune(buff[i])
				} else if buff[i] == 'O' {
					board.elements[coord] = '['
					coord = Coord{X: i*2 + 1, Y: j}
					board.elements[coord] = ']'
				}

			}
		}
		j++
	}
	board.maxY = j
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		instructions = instructions + buff
	}

	board.drawMapWithCoords(map[Coord]struct{}{botPosition: empty})

	for _, nextMovement := range instructions {
		// fmt.Println("Move", string(nextMovement))
		// fmt.Println("Bot position: ", botPosition)
		nextPosition := getNextPos(botPosition, nextMovement)
		// fmt.Println("Next position: ", nextPosition)

		if _, ok := board.elements[nextPosition]; !ok {
			botPosition = nextPosition
		} else if board.elements[nextPosition] == '[' || board.elements[nextPosition] == ']' {
			if board.moveRock(nextPosition, nextMovement) {
				botPosition = nextPosition
			}
		}

		// board.drawMapWithCoords(map[Coord]struct{}{botPosition: empty})
		// fmt.Println("---------------")

	}
	fmt.Println(board.calculateGPSSum())
}

func (board *Board) moveRock(rockPos Coord, direction rune) bool {
	switch direction {
	case '<':
		leftBox := Coord{X: rockPos.X - 1, Y: rockPos.Y}
		rightBox := rockPos
		leftBoxFuturePos := Coord{X: leftBox.X - 1, Y: rockPos.Y}
		if _, ok := board.elements[leftBoxFuturePos]; !ok {
			board.elements[leftBoxFuturePos] = '['
			board.elements[leftBox] = ']'
			delete(board.elements, rightBox)
			return true
		} else if board.elements[leftBoxFuturePos] == '#' {
			return false
		} else if board.elements[leftBoxFuturePos] == ']' {
			if board.moveRock(leftBoxFuturePos, direction) {
				board.elements[leftBoxFuturePos] = '['
				board.elements[leftBox] = ']'
				delete(board.elements, rightBox)
				return true
			} else {
				return false
			}
		}

	case '>':
		rightBox := Coord{X: rockPos.X + 1, Y: rockPos.Y}
		leftBox := rockPos
		rightBoxFuturePos := Coord{X: rightBox.X + 1, Y: rockPos.Y}
		if _, ok := board.elements[rightBoxFuturePos]; !ok {
			board.elements[rightBoxFuturePos] = ']'
			board.elements[rightBox] = '['
			delete(board.elements, leftBox)
			return true
		} else if board.elements[rightBoxFuturePos] == '#' {
			return false
		} else if board.elements[rightBoxFuturePos] == '[' {
			if board.moveRock(rightBoxFuturePos, direction) {
				board.elements[rightBoxFuturePos] = ']'
				board.elements[rightBox] = '['
				delete(board.elements, leftBox)
				return true
			} else {
				return false
			}
		}

	case '^':
		affectedPositions := board.obtainAllAffectedPositions(rockPos, direction)
		canMove := true
		for c := range affectedPositions {
			movedPosition := Coord{X: c.X, Y: c.Y - 1}
			if _, ok := affectedPositions[movedPosition]; ok {
				continue
			}
			if _, ok := board.elements[movedPosition]; ok {
				canMove = false
			}
		}
		if canMove {
			for c := range affectedPositions {
				delete(board.elements, c)
			}

			for c, r := range affectedPositions {
				board.elements[Coord{X: c.X, Y: c.Y - 1}] = r
			}
			return true
		} else {
			return false
		}

	case 'v':
		affectedPositions := board.obtainAllAffectedPositions(rockPos, direction)
		canMove := true
		for c := range affectedPositions {
			movedPosition := Coord{X: c.X, Y: c.Y + 1}
			if _, ok := affectedPositions[movedPosition]; ok {
				continue
			}
			if _, ok := board.elements[movedPosition]; ok {
				canMove = false
			}
		}
		if canMove {
			for c := range affectedPositions {
				delete(board.elements, c)
			}

			for c, r := range affectedPositions {
				board.elements[Coord{X: c.X, Y: c.Y + 1}] = r
			}
			return true
		} else {
			return false
		}

	}
	return false
}

func (board *Board) obtainAllAffectedPositions(rockPos Coord, direction rune) map[Coord]rune {
	ret := make(map[Coord]rune)
	var yMovement int
	var leftBox, rightBox, newLeftBox, newRightBox Coord

	if direction == '^' {
		yMovement = -1
	} else if direction == 'v' {
		yMovement = 1
	} else {
		panic("Wrong direction")
	}

	if board.elements[rockPos] == '[' {
		leftBox = rockPos
		rightBox = Coord{X: rockPos.X + 1, Y: rockPos.Y}
	} else if board.elements[rockPos] == ']' {
		leftBox = Coord{X: rockPos.X - 1, Y: rockPos.Y}
		rightBox = rockPos
	}
	ret[leftBox] = '['
	ret[rightBox] = ']'

	newLeftBox = Coord{X: leftBox.X, Y: leftBox.Y + yMovement}
	newRightBox = Coord{X: rightBox.X, Y: rightBox.Y + yMovement}
	if board.elements[newLeftBox] == '[' || board.elements[newLeftBox] == ']' {
		for c, r := range board.obtainAllAffectedPositions(newLeftBox, direction) {
			ret[c] = r
		}
	}

	if board.elements[newRightBox] == '[' || board.elements[newRightBox] == ']' {
		for c, r := range board.obtainAllAffectedPositions(newRightBox, direction) {
			ret[c] = r
		}
	}

	return ret
}

func getNextPos(coord Coord, movement rune) Coord {
	switch movement {
	case '^':
		return Coord{X: coord.X, Y: coord.Y - 1}
	case 'v':
		return Coord{X: coord.X, Y: coord.Y + 1}
	case '<':
		return Coord{X: coord.X - 1, Y: coord.Y}
	case '>':
		return Coord{X: coord.X + 1, Y: coord.Y}
	}
	return coord
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

func (board *Board) drawMapWithCoords(coords map[Coord]struct{}) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	minY := board.minY
	maxY := board.maxY

	for j := minY; j < maxY; j++ {
		for i := board.minX; i < board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			_, ok2 := coords[c]
			if ok && ok2 {
				fmt.Print(string(string(dot)))
			} else if ok && !ok2 {
				fmt.Print(string(dot))
			} else if !ok && ok2 {
				fmt.Print(string("@"))
			} else if !ok && !ok2 {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (board *Board) calculateGPSSum() int {
	ret := 0

	for c, r := range board.elements {
		if r == '[' {
			ret += c.X + (c.Y * 100)
		}
	}
	return ret
}
