package main

import (
	"bufio"
	"fmt"
	"os"
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

	score := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()

		// First my move
		switch string(buff[2]) {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		// Then the outcome
		switch buff {
		case "A X": // Stone vs Stone
			score += 3
		case "A Y": // Stone vs Paper
			score += 6
		case "A Z": // Stone vs Scissors
			score += 0
		case "B X": // Paper vs Stone
			score += 0
		case "B Y": // Paper vs Paper
			score += 3
		case "B Z": // Paper vs Scissors
			score += 6
		case "C X": // Scissors vs Stone
			score += 6
		case "C Y": // Scissors vs Paper
			score += 0
		case "C Z": // Scissors vs Scissors
			score += 3
		}
		// fmt.Printf("%s -> %d\n", string(buff), score)
	}

	fmt.Println(score)
}
