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
		case "X": // Lose
			score += 0
		case "Y": // Draw
			score += 3
		case "Z": // Win
			score += 6
		}

		// Then the outcome
		switch buff {
		case "A X": // Stone and lose -> Scissors
			score += 3
		case "A Y": // Stone and draw -> Stone
			score += 1
		case "A Z": // Stone and win -> Paper
			score += 2
		case "B X": // Paper and lose -> Stone
			score += 1
		case "B Y": // Paper and draw -> Paper
			score += 2
		case "B Z": // Paper and win -> Scissors
			score += 3
		case "C X": // Scissors and lose -> Paper
			score += 2
		case "C Y": // Scissors and draw -> Scissors
			score += 3
		case "C Z": // Scissors and win -> Stone
			score += 1
		}
		// fmt.Printf("%s -> %d\n", string(buff), score)
	}

	fmt.Println(score)
}
