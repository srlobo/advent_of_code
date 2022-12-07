package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type treeNode struct {
	parent   *treeNode
	children map[string]*treeNode
	size     int
	path     string
	t        string
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

	var tree, actual_node *treeNode
	tree = nil
	actual_node = nil
	for fileScanner.Scan() {
		buff := fileScanner.Text()

		fmt.Printf("CLI: %s\n", buff)
		if "$" == string(buff[0]) {
			// Command mode
			if "cd" == buff[2:4] {
				// Change dir
				new_path := buff[5:]
				if new_path == ".." {
					actual_node = actual_node.parent
				} else if new_path == "/" {
					// Root node
					new_dir := treeNode{path: new_path, size: 0, parent: actual_node, children: make(map[string]*treeNode), t: "d"}
					tree = &new_dir
					actual_node = &new_dir
				} else {
					new_dir := treeNode{path: new_path, size: 0, parent: actual_node, children: make(map[string]*treeNode), t: "d"}
					fmt.Printf("cd, new dir: %s\n", new_dir.path)
					actual_node.children[new_path] = &new_dir
					actual_node = &new_dir
				}
			}
		} else {
			// Inside of dir
			if "dir" != buff[0:4] {
				// File
				strSize := strings.Split(buff, " ")[0]
				path := strings.Split(buff, " ")[1]
				intSize, _ := strconv.Atoi(strSize)
				new_file := treeNode{path: path, size: intSize, parent: actual_node, children: nil, t: "f"}
				actual_node.children[path] = &new_file
			}

		}
	}
	fmt.Println("------------------------")
	tree.printTree(0)

	// Let's walk the tree breadth first and accumulating sizes upwards
	tree.calculateSizes()

	fmt.Println("------------------------")
	tree.printTree(0)

	fmt.Println("------------------------")
	fmt.Println(tree.totalSumBelow(100000))
}

func (tree treeNode) printTree(indent int) {
	padding := ""
	for i := 0; i < indent; i++ {
		padding += " "
	}
	if tree.t == "d" {
		fmt.Printf("%s- %s (dir, accsize: %d)\n", padding, tree.path, tree.size)
		for _, v := range tree.children {
			v.printTree(indent + 2)
		}
	} else {
		fmt.Printf("%s- %s (file, size=%d)\n", padding, tree.path, tree.size)
	}
}

func (tree *treeNode) calculateSizes() int {
	if tree.t == "d" {
		tree.size = 0
		for _, v := range tree.children {
			tree.size += v.calculateSizes()
		}
	}
	// fmt.Printf("path: %s, size: %d, type: %s\n", tree.path, tree.size, tree.t)
	return tree.size
}

func (tree treeNode) totalSumBelow(max int) int {
	ret := 0
	if tree.t == "d" {
		if tree.size <= max {
			ret += tree.size
		}

		for _, v := range tree.children {
			ret += v.totalSumBelow(max)
		}
	}
	return ret
}
