package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coords struct {
	X int
	Y int
}

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	headPos := Coords{X: 0, Y: 0}
	tailPos := Coords{X: 0, Y: 0}

	tailPosHistory := make(map[Coords]int)

	tailPosHistory[tailPos] = 1

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)

		direction := strings.Split(buff, " ")[0]
		amount, _ := strconv.Atoi(strings.Split(buff, " ")[1])

		for i := 0; i < amount; i++ {
			switch direction {
			case "U":
				headPos.X += 1
			case "D":
				headPos.X -= 1
			case "R":
				headPos.Y += 1
			case "L":
				headPos.Y -= 1
			}

			newPos := calculatePos(headPos, tailPos)
			if !equalCoords(newPos, tailPos) {
				tailPos = newPos
				tailPosHistory[tailPos] = 1
			}
		}
	}

	fmt.Println(len(tailPosHistory))
}

func equalCoords(a Coords, b Coords) bool {
	if a.X == b.X && a.Y == b.Y {
		return true
	}
	return false
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func reduceToOne(a int) int {
	if a < -1 {
		return -1
	} else if a > 1 {
		return 1
	}
	return a
}

func calculatePos(headPos Coords, tailPos Coords) Coords {
	newTailPos := tailPos

	diffX := headPos.X - tailPos.X
	diffY := headPos.Y - tailPos.Y

	if Abs(diffX)+Abs(diffY) > 1 && (Abs(diffX) > 1 || Abs(diffY) > 1) {
		newTailPos.X += reduceToOne(diffX)
		newTailPos.Y += reduceToOne(diffY)
	}

	fmt.Printf("headPos: (%d, %d), tailPos: (%d, %d); diffX: %d, diffy: %d;newTailPos: (%d, %d)\n", headPos.X, headPos.Y, tailPos.X, tailPos.Y, diffX, diffY, newTailPos.X, newTailPos.Y)

	return newTailPos
}
