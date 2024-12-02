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

	count := 0
	for fileScanner.Scan() {
		buff := strings.Split(fileScanner.Text(), " ")
		fmt.Println(buff)
		if checkLine(buff) {
			count++
		}
	}
	fmt.Println(count)
}

type Cmp func(a, b int) bool

func checkLine(line []string) bool {
	old_value, _ := strconv.Atoi(line[0])
	new_value, _ := strconv.Atoi(line[1])
	cmp := func(old_value, new_value int) Cmp {
		if old_value > new_value {
			return func(a, b int) bool {
				return a > b
			}
		} else {
			return func(a, b int) bool {
				return a < b
			}
		}
	}(old_value, new_value)

	for i := 1; i < len(line); i++ {
		new_value, _ := strconv.Atoi(line[i])
		if cmp(new_value, old_value) {
			fmt.Printf("new_value(%d) > old_value(%d)\n", new_value, old_value)
			return false
		}
		diff := abs(old_value - new_value)
		fmt.Printf("diff: %d\n", diff)
		if diff < 1 || diff > 3 {
			fmt.Printf("diff < 1 || diff > 3\n")
			return false
		}
		old_value = new_value
	}
	fmt.Println("return true")
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
