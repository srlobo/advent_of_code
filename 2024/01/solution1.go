package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"regexp"
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

	heapLeft := &IntHeap{}
	heapRight := &IntHeap{}

	heap.Init(heapLeft)
	heap.Init(heapRight)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		regex := regexp.MustCompile(" +")
		result := regex.Split(buff, -1)
		fmt.Println(result)
		left, _ := strconv.Atoi(result[0])
		heap.Push(heapLeft, left)
		right, _ := strconv.Atoi(result[1])
		heap.Push(heapRight, right)
		fmt.Println("left: %d, right: %d\n", left, right)
	}

	sum := 0
	for heapLeft.Len() > 0 {
		left := heap.Pop(heapLeft).(int)
		right := heap.Pop(heapRight).(int)
		sum += abs(left - right)
		fmt.Printf("left: %d, right: %d, sum: %d\n", left, right, sum)
	}
	fmt.Println(sum)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
