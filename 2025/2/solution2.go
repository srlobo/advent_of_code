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

	fileScanner.Scan()
	buff := fileScanner.Text()
	total := 0
	for val := range strings.SplitSeq(buff, ",") {
		tmp := strings.Split(val, "-")
		before, _ := strconv.Atoi(tmp[0])
		after, _ := strconv.Atoi(tmp[1])
		// fmt.Println(val, before, after)

		for i := before; i <= after; i++ {
			if check(i) {
				total += i
			}
		}
	}
	fmt.Println(total)
}

func check(number int) bool {
	strNumber := strconv.Itoa(number)
	var partial []byte

	var found bool
	for i := 0; i < len(strNumber)/2; i++ {
		partial = append(partial, strNumber[i])
		// fmt.Println("partial:", string(partial))
		if len(strNumber) % len(partial) != 0 {
			// fmt.Println("Not possible: ", len(strNumber), " % ", len(partial), "!= 0")
			continue
		}

		found = false
		for j := i + 1; j <= len(strNumber)-len(partial); j += len(partial) {

			// fmt.Println("Compare ", string(partial), strNumber[j:j+len(partial)])
			if strings.Compare(string(partial), strNumber[j:j+len(partial)]) == 0 {
				found = true
			} else {
				found = false
				break
			}
		}
		if found {
			break
		}
	}

	if found {
		fmt.Println(number, " is invalid")
	}
	return found
}
