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
	iterations := 100
	var total [4]int
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		parsedBuff := re.FindStringSubmatch(buff)
		// fmt.Println(parsedBuff)
		botX, _ := strconv.Atoi(parsedBuff[1])
		botY, _ := strconv.Atoi(parsedBuff[2])
		botvX, _ := strconv.Atoi(parsedBuff[3])
		botvY, _ := strconv.Atoi(parsedBuff[4])

		futurePosX := (botX + (botvX * iterations)) % board_maxX
		futurePosY := (botY + (botvY * iterations)) % board_maxY
		if futurePosX < 0 {
			futurePosX += board_maxX
		}
		if futurePosY < 0 {
			futurePosY += board_maxY
		}

		fmt.Printf("Bot (%d, %d) -> (%d, %d)\n", botX, botY, futurePosX, futurePosY)

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
	fmt.Println(total)

	fmt.Println(total[0] * total[1] * total[2] * total[3])
}
