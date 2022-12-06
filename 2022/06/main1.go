package main

import (
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
	readFile, err := os.ReadFile(filePath)

	check(err)

	buff := string(readFile)

	for i := 0; i < len(buff)-4; i++ {
		actual_packet := buff[i : i+4]
		fmt.Println(actual_packet)
		if !isAnyLetterFrequencyGreaterThanOne(actual_packet) {
			fmt.Println(i + 4)
			break
		}
	}

}

func isAnyLetterFrequencyGreaterThanOne(input string) bool {
	ret := make(map[string]int)
	for i := 0; i < len(input); i += 1 {
		char := string(input[i])
		fmt.Println(ret)
		if _, ok := ret[char]; ok {
			// key exists
			return true
		} else {
			ret[char] = 1
		}
	}

	return false
}
