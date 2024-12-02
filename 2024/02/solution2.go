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
		fmt.Print(buff)
		if checkLine(buff, 1) {
			fmt.Println(" Safe")
			count++
		} else {
			fmt.Println()
			for i := 0; i < len(buff); i++ {
				new_buff := make([]string, len(buff)-1)
				c := 0
				for j := 0; j < len(buff); j++ {
					if i != j {
						new_buff[c] = buff[j]
						c++
					}
				}

				fmt.Printf("%v + %v -> %v", buff[:i], buff[i+1:], new_buff)
				if checkLine(new_buff, 0) {
					fmt.Println(" Safe (second try)")
					count++
					break
				} else {
					fmt.Println(" UnSafe (second try)")
				}
			}
		}
	}
	fmt.Println(count)
}

func checkLine(line []string, max_errors int) bool {
	old_value, _ := strconv.Atoi(line[0])
	new_value, _ := strconv.Atoi(line[1])

	old_diff := old_value - new_value
	for i := 1; i < len(line); i++ {
		new_value, _ := strconv.Atoi(line[i])
		if old_diff*(old_value-new_value) < 0 {
			// fmt.Printf("new_value(%d) > old_value(%d)\n", new_value, old_value)
			return false
		}
		diff := abs(old_value - new_value)
		// fmt.Printf("diff: %d\n", diff)
		if diff < 1 || diff > 3 {
			// fmt.Printf("diff < 1 || diff > 3\n")
			return false
		}
		old_diff = old_value - new_value
		old_value = new_value
	}
	// fmt.Println("return true")
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
