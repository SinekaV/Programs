//Create a binary search tree and find the location of the element in the level starting from 0 upto 16 levels and also search the time complexity for searching that particular element
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value}
	}
	if value < root.Value {
		root.Left = insertNode(root.Left, value)
	} else {
		root.Right = insertNode(root.Right, value)
	}
	return root
}

func findElementLevel(root *TreeNode, value int) int {
	level := 0
	current := root
	for current != nil {
		if value == current.Value {
			return level
		} else if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
		level++
	}
	return -1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var root *TreeNode
	for i := 0; i < 100; i++ {
		root = insertNode(root, rand.Intn(1000))
	}

	elementToSearch := 42
	level := findElementLevel(root, elementToSearch)

	if level >= 0 {
		fmt.Printf("Element %d found at level %d\n", elementToSearch, level)
	} else {
		fmt.Printf("Element %d not found in the tree\n", elementToSearch)
	}
}
