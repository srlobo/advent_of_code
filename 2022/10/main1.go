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

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cycle := 0
	waiting := false
	payload := 0
	registerX := 1
	result := 0
	current_instruction := ""
	for {
		cycle += 1

		fmt.Printf("Cycle: %d, instruction: %s, payload: %d; X: %d\n", cycle, current_instruction, payload, registerX)
		if (cycle-20)%40 == 0 {
			fmt.Printf("TOTAL: Cycle: %d, X: %d\n", cycle, registerX)
			result += cycle * registerX
		}

		if !waiting {
			if !fileScanner.Scan() {
				fmt.Println("EOF")
				break
			}
			buff := fileScanner.Text()

			splited_buff := strings.Split(buff, " ")
			current_instruction = splited_buff[0]
			switch current_instruction {
			case "noop":
			case "addx":
				payload, _ = strconv.Atoi(splited_buff[1])
				waiting = true
			}
		} else {
			registerX += payload
			waiting = false
		}

	}
	fmt.Println(result)

}
