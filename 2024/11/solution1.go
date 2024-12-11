package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var empty = struct{}{}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
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

	total := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		for _, s := range strings.Split(buff, " ") {
			fmt.Println(s)
			total += process(s, 1)
		}
	}
	fmt.Println(total)
}

func process(stone string, depth int) int {
	if depth > 25 {
		return 1
	}

	if stone == "0" {
		// fmt.Println("0 -> 1; depth", depth)
		return process("1", depth+1)
	}

	size := len(stone)
	if size%2 == 0 {
		firstHalf, _ := strconv.Atoi(stone[:size/2])
		// fmt.Println(stone, "->", firstHalf, "; depth", depth)
		total := process(strconv.Itoa(firstHalf), depth+1)
		secondHalf, _ := strconv.Atoi(stone[size/2:])
		// fmt.Println(stone, "->", secondHalf, "; depth", depth)
		total += process(strconv.Itoa(secondHalf), depth+1)
		return total
	} else {
		num, _ := strconv.Atoi(stone)
		num = num * 2024
		// fmt.Println(stone, "->", num, "; depth", depth)
		return process(strconv.Itoa(num), depth+1)
	}
}
