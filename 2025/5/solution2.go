package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var empty = struct{}{}

type idRange struct {
	lower int
	upper int
}

type idRanges map[idRange]struct{}

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

	freshIngredientsIDRanges := make(idRanges)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if buff == "" {
			break
		}
		numbers := strings.Split(buff, "-")
		lower, _ := strconv.Atoi(numbers[0])
		upper, _ := strconv.Atoi(numbers[1])

		r := idRange{lower, upper}
		freshIngredientsIDRanges[r] = empty

	}

	freshIngredientsIDRanges.removeIntersections()

	fmt.Println(freshIngredientsIDRanges.count())

}

func printIDRange(ingredientsIDRanges idRanges) {
	for r, _ := range ingredientsIDRanges {
		fmt.Println(r.lower, "-", r.upper)
	}
}

func (idranges *idRanges) count() int {
	count := 0
	for r, _ := range *idranges {
		count += r.upper - r.lower + 1
	}
	return count

}

func haveIntersection(idrange1, idrange2 idRange) bool {
	fmt.Printf("Testing %v vs %v\n", idrange1, idrange2)
	if idrange2.lower >= idrange1.lower && idrange2.lower <= idrange1.upper {
	fmt.Println("yes")
		return true
	}
	if idrange2.upper >= idrange1.lower && idrange2.upper <= idrange1.upper {
	fmt.Println("yes")
		return true
	}
	if idrange1.lower >= idrange2.lower && idrange1.lower <= idrange2.upper {
	fmt.Println("yes")
		return true
	}
	if idrange1.upper >= idrange2.lower && idrange1.upper <= idrange2.upper {
	fmt.Println("yes")
		return true
	}

	fmt.Println("no")
	return false
}

func fusionate(idrange1, idrange2 idRange) idRange {
	fmt.Printf("Join %v vs %v\n", idrange1, idrange2)

	var lower, upper int
	if idrange1.lower < idrange2.lower {
		lower = idrange1.lower
	} else {
		lower = idrange2.lower
	}

	if idrange1.upper > idrange2.upper {
		upper = idrange1.upper
	} else {
		upper = idrange2.upper
	}

	fmt.Println("res", idRange{lower, upper})
	return idRange{lower, upper}
}

func (idranges *idRanges) removeIntersections() {
	for {
		changes := false
		for r1, _ := range *idranges {
			for r2, _ := range *idranges {
				if r1 == r2 {
					continue
				}
				if haveIntersection(r1, r2) {
					newRange := fusionate(r1, r2)
					delete(*idranges, r1)
					delete(*idranges, r2)
					(*idranges)[newRange] = empty
					changes = true
					break
				}
				if changes {
					break
				}
			}
			if changes {
				break
			}
		}
		if !changes {
			break
		}
	}
}
