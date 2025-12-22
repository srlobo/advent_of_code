package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	fileScanner.Scan()
	buff := fileScanner.Text()
	total := 0
	for val := range strings.SplitSeq(buff, ",") {
		tmp := strings.Split(val, "-")
		before, _ := strconv.Atoi(tmp[0])
		after, _ := strconv.Atoi(tmp[1])
		// fmt.Println(val, before, after)

		for i := before; i <= after; i++ {
			if checkOnlyTwo(i) {
				total += i
			}
		}
	}
	fmt.Println(total)
}

func checkOnlyTwo(number int) bool {
	strNumber := strconv.Itoa(number)
	if len(strNumber)%2 == 1 {
		return false
	}

	for i := 0; i < len(strNumber)/2; i++ {
		if strNumber[i] != strNumber[len(strNumber)/2+i] {
			return false
		}
	}
	return true
}
