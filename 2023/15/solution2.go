package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var empty = struct{}{}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

type Lens struct {
	focal_length int
	label        string
}
type (
	Box   []Lens
	Boxes map[int]Box
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

	boxes := make(Boxes)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		for _, instruction := range strings.Split(buff, ",") {
			boxes.processInstruction(instruction)
			// boxes.print()
			// fmt.Println()
		}
	}
	fmt.Println(boxes.calculateFocusingPower())
}

func (boxes *Boxes) print() {
	for box_n, box := range *boxes {
		fmt.Printf("Box %d: ", box_n)
		for _, lens := range box {
			fmt.Printf("[%s %d] ", lens.label, lens.focal_length)
		}
		fmt.Println()
	}
}

func (boxes *Boxes) processInstruction(instruction string) {
	fmt.Println(instruction)

	r := regexp.MustCompile(`([a-z]+)([=-])([0-9]*)`)
	expression := r.FindStringSubmatch(instruction)
	label := expression[1]
	operator := expression[2]
	focal_length, _ := strconv.Atoi(expression[3])
	box_n := calculateHash(label)
	// fmt.Println("label:", label, "operator:", operator, "focal_length:", focal_length, "box_n:", box_n)

	if operator == "-" {
		if _, ok := (*boxes)[box_n]; ok {
			for i := 0; i < len((*boxes)[box_n]); i++ {
				lens := (*boxes)[box_n][i]
				if lens.label == label {
					(*boxes)[box_n] = append((*boxes)[box_n][:i], (*boxes)[box_n][i+1:]...)
					break
				}
			}
			if len((*boxes)[box_n]) == 0 {
				delete(*boxes, box_n)
			}
		}
	} else if operator == "=" {
		if _, ok := (*boxes)[box_n]; !ok {
			box := make(Box, 0)
			box = append(box, Lens{focal_length: focal_length, label: label})
			(*boxes)[box_n] = box
		} else {
			found := false
			for i := 0; i < len((*boxes)[box_n]); i++ {
				lens := (*boxes)[box_n][i]
				if lens.label == label {
					(*boxes)[box_n][i] = Lens{focal_length: focal_length, label: label}
					found = true
				}
			}
			if !found {
				(*boxes)[box_n] = append((*boxes)[box_n], Lens{focal_length: focal_length, label: label})
			}
		}
	}
}

func (boxes *Boxes) calculateFocusingPower() int {
	total := 0
	for box_n, box := range *boxes {
		for slot, lens := range box {
			box_total := (1 + box_n) * (1 + slot) * lens.focal_length
			// fmt.Println("box_n:", box_n, "slot:", slot, "lens.focal_length:", lens.focal_length, "box_total:", box_total)
			total += box_total
		}
	}

	return total
}

func calculateHash(input string) int {
	ret := 0
	for _, c := range input {
		ret += int(c)
		ret *= 17
		ret %= 256
	}
	return ret
}
