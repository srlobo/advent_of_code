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

	for {
		if !fileScanner.Scan() {
			break
		}
		buff := fileScanner.Text()
		set1 := letterFrequency(buff)

		fileScanner.Scan()
		buff = fileScanner.Text()
		set2 := letterFrequency(buff)

		set1Iset2 := setIntersection(set1, set2)

		fileScanner.Scan()
		buff = fileScanner.Text()
		set3 := letterFrequency(buff)

		fullIntesection := setIntersection(set1Iset2, set3)

		fmt.Println(fullIntesection)
		for k := range fullIntesection {
			total += getPriority(k)
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

func setIntersection(a map[string]int, b map[string]int) map[string]int {
	ret := make(map[string]int)
	for k := range a {
		if _, ok := b[k]; ok {
			ret[k] = 1
		}
	}

	return ret
}
