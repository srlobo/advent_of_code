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

	candidates := make(map[[2]int]int)
	for j := 0; j < len(treeGrid); j++ {
		row_max := -1
		for i := 0; i < len(treeGrid[j]); i++ {
			coords := [2]int{i, j}
			if treeGrid[i][j] > row_max {
				candidates[coords] = 1
				row_max = treeGrid[i][j]
			}
		}

		row_max = -1
		for i := len(treeGrid[j]) - 1; i >= 0; i-- {
			coords := [2]int{i, j}
			if treeGrid[i][j] > row_max {
				candidates[coords] = 1
				row_max = treeGrid[i][j]
			}
		}
	}

	for i := 0; i < len(treeGrid[0]); i++ {
		col_max := -1
		for j := 0; j < len(treeGrid); j++ {
			coords := [2]int{i, j}
			if treeGrid[i][j] > col_max {
				candidates[coords] = 1
				col_max = treeGrid[i][j]
			}
		}

		col_max = -1
		for j := len(treeGrid) - 1; j >= 0; j-- {
			coords := [2]int{i, j}
			if treeGrid[i][j] > col_max {
				candidates[coords] = 1
				col_max = treeGrid[i][j]
			}
		}
	}

	printTreeGrid(treeGrid, candidates)
	fmt.Println(len(candidates))
}

func printTreeGrid(treeGrid [][]int, candidates map[[2]int]int) {
	for j := 0; j < len(treeGrid); j++ {
		for i := 0; i < len(treeGrid[j]); i++ {
			coords := [2]int{i, j}

			if _, ok := candidates[coords]; ok {
				color(treeGrid[i][j])
			} else {
				fmt.Print(treeGrid[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func color(i int) {
	colored := fmt.Sprintf("\x1b[%dm%d\x1b[0m", 34, i)
	fmt.Print(colored)
}
