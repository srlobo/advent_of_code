package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var coords []Coord
	for fileScanner.Scan() {
		buff := strings.Split(fileScanner.Text(), ",")
		x, _ := strconv.Atoi(buff[0])
		y, _ := strconv.Atoi(buff[1])
		coords = append(coords, Coord{x, y})
	}
	fmt.Println(coords)
	maxArea := 0
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			currentArea := area(coords[i], coords[j])
			if currentArea > maxArea {
				maxArea = currentArea

			}
		}
	}
	fmt.Println(maxArea)
}

type Coord struct {
	X int
	Y int
}

func area(a, b Coord) int {
	return (abs(b.X-a.X) + 1) * (abs(b.Y-a.Y) + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
