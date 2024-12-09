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

	fileScanner.Scan()
	buff := fileScanner.Text()
	id := 0
	disk := Node{-1, false, 0, nil, nil}
	visiting_node := &disk
	for i := 0; i < len(buff); i++ {
		num, _ := strconv.Atoi(string(buff[i]))
		// position is even, it's a block.
		if i%2 == 0 {
			id = i / 2
			visiting_node.append(&Node{id, true, num, nil, nil})
			visiting_node = visiting_node.next
		} else { // Position is odd, it's space, let's fill with the latest numbers
			visiting_node.append(&Node{0, false, num, nil, nil})
			visiting_node = visiting_node.next
		}
	}

	// disk.printDisk()
	visiting_node = &disk
	for visiting_node.next != nil {
		visiting_node = visiting_node.next
	}

	for true {
		if visiting_node.prev == nil {
			break
		}

		// Is it a space?
		if !visiting_node.is_file {
			visiting_node = visiting_node.prev
			continue
		}
		visiting_space := &disk
		for true {
			if visiting_space.next == nil {
				break
			} else if visiting_space == visiting_node {
				break
			}
			if visiting_space.is_file {
				visiting_space = visiting_space.next
				continue
			}
			if visiting_space.size >= visiting_node.size {
				node := &Node{visiting_node.id, visiting_node.is_file, visiting_node.size, nil, nil}
				visiting_space.prev.insert(node)
				if visiting_space.size > node.size {
					visiting_space.size -= node.size
				} else {
					visiting_space.delete()
				}
				visiting_node.is_file = false
				break

			}
			visiting_space = visiting_space.next
		}
		visiting_node = visiting_node.prev
		// disk.printDisk()
	}
	fmt.Println("Total: ", disk.calculateChecksum())
}

func (node *Node) calculateChecksum() int {
	total := 0
	n := node
	pos := 0
	for true {
		if n.is_file {
			for i := 0; i < n.size; i++ {
				total = total + pos*n.id
				pos++
			}
		} else {
			for i := 0; i < n.size; i++ {
				pos++
			}
		}
		if n.next == nil {
			break
		}
		n = n.next
	}
	return total
}

func (node *Node) printDisk() {
	n := node
	for true {
		if n.is_file {
			for i := 0; i < n.size; i++ {
				fmt.Print(n.id)
			}
		} else {
			for i := 0; i < n.size; i++ {
				fmt.Print(".")
			}
		}
		if n.next == nil {
			break
		}
		n = n.next
	}
	fmt.Println()
}

func (node *Node) append(new_node *Node) {
	node.next = new_node
	new_node.prev = node
}

func (node *Node) insert(new_node *Node) {
	if node.next == nil {
		node.append(new_node)
	} else {
		left := node
		right := node.next

		left.next = new_node
		new_node.prev = left
		new_node.next = right
		right.prev = new_node
	}
}

func (node *Node) delete() {
	left := node.prev
	right := node.next

	if left != nil {
		left.next = right
	}
	if right != nil {
		right.prev = left
	}
	node.next = nil
	node.prev = nil
}

type Node struct {
	id      int
	is_file bool
	size    int
	next    *Node
	prev    *Node
}
