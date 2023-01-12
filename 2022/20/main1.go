package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Vector []int

type Node struct {
	previous *Node
	next     *Node
	val      int
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

	var nums Vector
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		num, _ := strconv.Atoi(buff)
		nums = append(nums, num)
	}

	var first_node, pointer *Node
	node := Node{val: nums[0]}
	first_node = &node
	pointer = first_node

	for i := 1; i < len(nums); i++ {
		node := Node{previous: pointer, val: nums[i]}
		pointer.next = &node
		pointer = &node
	}

	pointer.next = first_node
	first_node.previous = pointer

	first_node.PrintLinkedList()

	for j := 0; j < len(nums); j++ {
		fmt.Printf("Move %d; len: %d\n", nums[j], len(nums))
		first_node.Move(nums[j], len(nums))
		first_node.PrintLinkedList()
		//fmt.Println()
	}

	first_node.PrintLinkedList()

	ret := 0

	pointer = first_node.Find(0)
	for i := 0; i < 1000%len(nums); i++ {
		pointer = pointer.next
	}
	fmt.Println(pointer.val)
	ret += pointer.val

	pointer = first_node.Find(0)
	for i := 0; i < 2000%len(nums); i++ {
		pointer = pointer.next
	}
	fmt.Println(pointer.val)
	ret += pointer.val

	pointer = first_node.Find(0)
	for i := 0; i < 3000%len(nums); i++ {
		pointer = pointer.next
	}
	fmt.Println(pointer.val)
	ret += pointer.val
	fmt.Println(ret)

}

func (node *Node) PrintLinkedList() {
	var v []int
	var nodes []*Node
	p := node
	for {
		v = append(v, p.val)
		nodes = append(nodes, p)
		p = p.next
		if p == node {
			break
		}

	}
	// for _, node := range nodes {
	//	fmt.Printf("%p -> %v\n", node, node)
	//	fmt.Println(nodes)
	//}

	fmt.Println(v)

}

func (node *Node) Move(element, size int) {
	// node.PrintLinkedList()
	if element == 0 {
		return
	}
	// First we find origin
	var p, origin *Node
	p = node
	for {
		if p.val == element {
			break
		}
		p = p.next
	}

	origin = p
	element = element % size

	// fmt.Printf("Origin: %p; %v\n", origin, *origin)
	// fmt.Printf("Advance %d elements\n", element)

	// Now let's advance
	if element > 0 {
		for i := 0; i < element%size; i++ {
			p = p.next
		}
	} else if element < 0 {
		for i := 0; i >= element%size; i-- {
			p = p.previous
		}
	}
	// fmt.Printf("Destination: %p; %v\n", p, *p)
	// Remove the origin element
	var previous, next *Node
	previous = origin.previous
	next = origin.next
	previous.next = next
	next.previous = previous
	// fmt.Println("After origin")

	// And then insert into current position
	previous = p
	next = p.next

	previous.next = origin
	next.previous = origin
	origin.next = next
	origin.previous = previous

	// fmt.Println("After destination")
	// node.PrintLinkedList()
}

func (node *Node) Find(element int) *Node {
	var p *Node
	p = node
	for {
		if p.val == element {
			return p
		}
		p = p.next
	}
}
