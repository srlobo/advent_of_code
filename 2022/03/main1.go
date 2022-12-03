package main

import (
	"bufio"
	"fmt"
	"os"
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

	total := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()

		// fmt.Println(buff)

		bufLength := len(buff)
		firstHalf := buff[0 : bufLength/2]
		secondtHalf := buff[bufLength/2 : bufLength]

		// fmt.Println(firstHalf)
		firsHalfFrequency := letterFrequency(firstHalf)

		// fmt.Println(firsHalfFrequency)

		// fmt.Println(secondtHalf)
		secondHalfFrequency := letterFrequency(secondtHalf)
		// fmt.Println(secondHalfFrequency)

		for k := range firsHalfFrequency {
			if _, ok := secondHalfFrequency[k]; ok {
				// fmt.Printf("Found, letter %s, priority %d\n", k, getPriority(k))
				total += getPriority(k)
			}
		}
	}

	fmt.Println(total)
}

func letterFrequency(input string) map[string]int {
	ret := make(map[string]int)
	for i := 0; i < len(input); i += 1 {
		char := string(input[i])
		if _, ok := ret[char]; ok {
			// key exists
			ret[char] += 1
		} else {
			ret[char] = 1
		}
	}

	return ret
}

func getPriority(input string) int {
	if strings.ToLower(input) == input {
		return int([]byte(input)[0]-[]byte("a")[0]) + 1
	} else {
		return int([]byte(input)[0]-[]byte("A")[0]) + 27
	}
}
