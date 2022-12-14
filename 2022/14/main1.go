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

type Coord struct {
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

	mapa := make(map[Coord]rune)
	mapa[Coord{X: 500, Y: 0}] = '+'
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		extremes := strings.Split(buff, " -> ")
		for i := 1; i < len(extremes); i++ {
			left := CoordFromText(extremes[i-1])
			right := CoordFromText(extremes[i])
			if left.X == right.X {
				extreme1, extreme2 := getSortedYExtremes(left, right)
				fmt.Printf("Beam from %d to %d, X: %d\n", extreme1, extreme2, left.X)
				for j := extreme1; j <= extreme2; j++ {
					fmt.Printf("Adding point (%d to %d)\n", left.X, j)
					mapa[Coord{X: left.X, Y: j}] = 'b'
				}
			} else {
				extreme1, extreme2 := getSortedXExtremes(left, right)
				fmt.Printf("Beam from %d to %d, Y: %d\n", extreme1, extreme2, left.Y)
				for j := extreme1; j <= extreme2; j++ {
					mapa[Coord{X: j, Y: left.Y}] = 'b'
				}
			}
		}
	}
	drawMap(mapa)

	_, _, _, maxY := getBoundingBox(mapa)

	// starting with the sand
	end := false
	counter := 0
	for {
		counter += 1
		s := Coord{X: 500, Y: 0}
		for {
			tentative_s := Coord{X: s.X, Y: s.Y + 1}
			// fmt.Printf("Trying (%d, %d)\n", tentative_s.X, tentative_s.Y)
			moved := false
			if _, ok := mapa[tentative_s]; !ok { // There's nothing underneath
				s = tentative_s
				moved = true
			} else {
				tentative_s = Coord{X: s.X - 1, Y: s.Y + 1}
				// fmt.Printf("Trying (%d, %d)\n", tentative_s.X, tentative_s.Y)
				if _, ok := mapa[tentative_s]; !ok { // There's nothing underneath on the left
					s = tentative_s
					moved = true
				} else {
					tentative_s = Coord{X: s.X + 1, Y: s.Y + 1}
					// fmt.Printf("Trying (%d, %d)\n", tentative_s.X, tentative_s.Y)
					if _, ok := mapa[tentative_s]; !ok { // There's nothing underneath on the right
						s = tentative_s
						moved = true
					}
				}
			}
			if s.Y == maxY {
				end = true
			}

			if !moved || end {
				break
			}
		}
		mapa[s] = 'o'
		// drawMap(mapa)
		// fmt.Println()
		if end {
			break
		}
	}
	drawMap(mapa)
	fmt.Println(counter - 1)
}

func CoordFromText(buff string) Coord {
	point := strings.Split(buff, ",")

	x, _ := strconv.Atoi(point[0])
	y, _ := strconv.Atoi(point[1])

	return Coord{X: x, Y: y}
}

func getSortedYExtremes(left, right Coord) (int, int) {
	if left.Y < right.Y {
		return left.Y, right.Y
	} else {
		return right.Y, left.Y
	}
}

func getSortedXExtremes(left, right Coord) (int, int) {
	if left.X < right.X {
		return left.X, right.X
	} else {
		return right.X, left.X
	}
}

func getBoundingBox(mapa map[Coord]rune) (int, int, int, int) {
	var minX, minY, maxX, maxY int
	// minY is always 0
	for coord := range mapa {
		// fmt.Printf("Comparing (%d, %d)\n", coord.X, coord.Y)
		if coord.X < minX || minX == 0 {
			minX = coord.X
		} else if coord.X > maxX || maxX == 0 {
			maxX = coord.X
		}

		if coord.Y > maxY || maxY == 0 {
			maxY = coord.Y
		}
	}

	return minX, minY, maxX, maxY

}
func drawMap(mapa map[Coord]rune) {
	minX, minY, maxX, maxY := getBoundingBox(mapa)

	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", minX, minY, maxX, maxY)

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			dot, ok := mapa[Coord{X: i, Y: j}]
			if ok {
				fmt.Print(string(dot))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
