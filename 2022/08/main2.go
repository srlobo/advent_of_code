package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

	var treeGrid [][]int

	for fileScanner.Scan() {
		buff := fileScanner.Text()

		var row []int
		for i := 0; i < len(buff); i++ {
			val, _ := strconv.Atoi(string(buff[i]))
			row = append(row, val)
		}
		treeGrid = append(treeGrid, row)
	}

	// printTreeGrid(treeGrid)
	// fmt.Println(calculateScenicScore(treeGrid, [2]int{3, 2}))

	maxScenicScore := 0

	for j := 0; j < len(treeGrid); j++ {
		for i := 0; i < len(treeGrid[j]); i++ {
			coords := [2]int{i, j}
			currentScenicScore := calculateScenicScore(treeGrid, coords)
			// fmt.Printf("(%d, %d), scenic: %d\n", i, j, currentScenicScore)
			if currentScenicScore > maxScenicScore {
				maxScenicScore = currentScenicScore
			}
		}
	}

	fmt.Println(maxScenicScore)
}

func calculateScenicScore(treeGrid [][]int, position [2]int) int {
	x := position[0]
	y := position[1]
	currentHeight := treeGrid[x][y]
	// fmt.Printf("Current height: %d\n", currentHeight)
	score := 1
	var distance int

	// Up
	distance = 0
	for i := x - 1; i >= 0; i-- {
		// fmt.Printf("Comparing with (%d, %d): %d vs %d\n", i, y, currentHeight, treeGrid[i][y])
		distance++
		if currentHeight <= treeGrid[i][y] {
			// fmt.Println("Salimos")
			break
		}
	}
	// fmt.Printf("Up len: %d\n", distance)
	score *= distance

	// Left
	distance = 0
	for j := y - 1; j >= 0; j-- {
		// fmt.Printf("Comparing with (%d, %d): %d vs %d; d: %d\n", x, j, currentHeight, treeGrid[x][j], distance)
		distance++
		if currentHeight <= treeGrid[x][j] {
			// fmt.Println("Salimos")
			break
		}
	}
	// fmt.Printf("Left len: %d\n", distance)
	score *= distance

	// Down
	distance = 0
	for i := x + 1; i < len(treeGrid[y]); i++ {
		// fmt.Printf("Comparing with (%d, %d): %d vs %d\n", i, y, currentHeight, treeGrid[i][y])
		distance++
		if currentHeight <= treeGrid[i][y] {
			// fmt.Println("Salimos")
			break
		}
	}
	// fmt.Printf("Down len: %d\n", distance)
	score *= distance

	// Right
	distance = 0
	for j := y + 1; j < len(treeGrid); j++ {
		// fmt.Printf("Comparing with (%d, %d): %d vs %d\n", x, j, currentHeight, treeGrid[x][j])
		distance++
		if currentHeight <= treeGrid[x][j] {
			// fmt.Println("Salimos")
			break
		}
	}
	// fmt.Printf("Right len: %d\n", distance)
	score *= distance

	return score
}

func printTreeGrid(treeGrid [][]int) {
	for j := 0; j < len(treeGrid); j++ {
		for i := 0; i < len(treeGrid[j]); i++ {
			fmt.Print(treeGrid[j][i])
		}
		fmt.Print("\n")
	}
}
