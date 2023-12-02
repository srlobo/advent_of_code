package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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
		// fmt.Println("=====================================")
		// fmt.Println(buff)
		i := 0
		first_number := 0
		last_number := 0
		for {
			if i > (len(buff)-1) || first_number != 0 {
				break
			}

			//	fmt.Println()
			//	fmt.Println(buff[i:])
			// fmt.Println(num_array)
			changed := false
			for n, number := range numbers {
				if strings.HasPrefix(buff[i:], number) {
					first_number = n + 1
					changed = true
					break
				}
			}
			if changed {
				break
			}

			// fmt.Printf("buff[i] -> %d\n", buff[i])
			// fmt.Printf("buff[i] - '0' -> %d\n", buff[i]-'0')
			if (int(buff[i]) - '0') < 10 {
				first_number = int(buff[i]) - '0'
				break
			}
			i += 1
		}

		i = len(buff) - 1
		for {
			if i < 0 || last_number != 0 {
				break
			}

			//	fmt.Println()
			//	fmt.Println(buff[i:])
			// fmt.Println(num_array)
			changed := false
			for n, number := range numbers {
				if strings.HasPrefix(buff[i:], number) {
					last_number = n + 1
					changed = true
					break
				}
			}
			if changed {
				break
			}

			// fmt.Printf("buff[i] -> %d\n", buff[i])
			// fmt.Printf("buff[i] - '0' -> %d\n", buff[i]-'0')
			if (int(buff[i]) - '0') < 10 {
				last_number = int(buff[i]) - '0'
				break
			}
			i = i - 1
		}
		partial, _ := strconv.Atoi(fmt.Sprintf("%d%d", first_number, last_number))
		fmt.Printf("%s -> %d\n", buff, partial)
		total += partial
	}
	fmt.Println(total)
}
