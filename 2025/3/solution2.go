package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		fmt.Println(buff)
		battery := Battery{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		position := 0
		for batteryPos := 0; batteryPos < len(battery); batteryPos++ {
			for j := position; j <= len(buff)-(len(battery)-batteryPos); j++ {
				current, _ := strconv.Atoi(string(buff[j]))
				if battery[batteryPos] < current {
					position = j + 1
					battery[batteryPos] = current
				}

			}
		}
		fmt.Println(battery.total())

		total += battery.total()
	}
	fmt.Println(total)
}

type Battery [12]int

func (battery *Battery) reset(position int) {
	for i := position; i < len(battery); i++ {
		battery[position] = 0
	}
}

func (battery *Battery) total() int {
	total := 0
	exp := 1
	for i := len(battery) - 1; i >= 0; i-- {
		total += battery[i] * exp
		exp = exp * 10
	}
	return total
}
