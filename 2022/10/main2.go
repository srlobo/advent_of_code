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
	current_instruction := ""
	var current_line [40]bool
	var lines [6][40]bool

	for {
		cycle += 1
		line_pos := (cycle - 1) % 40

		fmt.Printf("Cycle: %d, instruction: %s, payload: %d; X: %d\n", cycle, current_instruction, payload, registerX)
		printLine(current_line)
		if line_pos == 0 && cycle > 1 {
			lines[(cycle-1)/40-1] = current_line
		}

		// Calculate if the CRT have to light the pixel
		current_line[line_pos] = Abs(line_pos-registerX) <= 1

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
	printLines(lines)
}

func printLine(line [40]bool) {
	for j := 0; j < len(line); j++ {
		if line[j] {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func printLines(lines [6][40]bool) {
	for i := 0; i < len(lines); i++ {
		printLine(lines[i])
	}
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
