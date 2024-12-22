package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	sum := uint64(0)
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		seed, _ := strconv.ParseUint(buff, 10, 64)

		fmt.Print(seed)
		fmt.Print(": ")

		for i := 0; i < 2000; i++ {
			seed = getNext(seed)
		}
		fmt.Println(seed)
		sum += seed

	}
	fmt.Println(sum)
}

func getNext(actual uint64) uint64 {
	next := actual
	next = mix(next*64, next)
	next = prune(next)

	next = mix(next/32, next)
	next = prune(next)

	next = mix(2048*next, next)
	next = prune(next)

	return next
}

func mix(a, b uint64) uint64 {
	return a ^ b
}

func prune(a uint64) uint64 {
	return a % 16777216
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
