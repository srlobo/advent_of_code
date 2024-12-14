package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	board_maxX := 101
	board_maxY := 103
	var bots []BotVelocity
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		// fmt.Println(buff)
		parsedBuff := re.FindStringSubmatch(buff)
		// fmt.Println(parsedBuff)
		botX, _ := strconv.Atoi(parsedBuff[1])
		botY, _ := strconv.Atoi(parsedBuff[2])
		botvX, _ := strconv.Atoi(parsedBuff[3])
		botvY, _ := strconv.Atoi(parsedBuff[4])

		bot := BotVelocity{botX, botY, botvX, botvY}
		bots = append(bots, bot)
	}
	fmt.Println("Initial state")

	iterations := 0
	for {
		total := [4]int{0, 0, 0, 0}
		var newBots []BotVelocity
		for _, c := range bots {
			futurePosX := (c.X + (c.vX)) % board_maxX
			futurePosY := (c.Y + (c.vY)) % board_maxY
			if futurePosX < 0 {
				futurePosX += board_maxX
			}
			if futurePosY < 0 {
				futurePosY += board_maxY
			}

			newBots = append(newBots, BotVelocity{futurePosX, futurePosY, c.vX, c.vY})
			if futurePosX < board_maxX/2 && futurePosY < board_maxY/2 {
				total[0]++
			} else if futurePosX > board_maxX/2 && futurePosY < board_maxY/2 {
				total[2]++
			} else if futurePosX < board_maxX/2 && futurePosY > board_maxY/2 {
				total[1]++
			} else if futurePosX > board_maxX/2 && futurePosY > board_maxY/2 {
				total[3]++
			}

		}
		bots = newBots
		if iterations%103 == 32 || iterations%101 == 83 {
			fmt.Println("Iteration: ", iterations, "len", len(bots))
			board := Board{make(BoardElements), 0, 0, board_maxX, board_maxY}
			for _, c := range bots {
				board.elements[Coord{c.X, c.Y}] = '#'
			}
			board.drawMap()
		}

		iterations++
	}
}

type Coord struct {
	X int
	Y int
}

type BotVelocity struct {
	X  int
	Y  int
	vX int
	vY int
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
	fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)

	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			_, ok := board.elements[c]
			if ok {
				fmt.Print(string("#"))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
