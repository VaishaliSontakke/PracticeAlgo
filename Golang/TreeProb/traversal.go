package main

import (
	"fmt"
)

func main() {
	node1 := newTreeNode(5)
	node1.left = newTreeNode(2)
	node1.right = newTreeNode(7)
	node1.right.left = newTreeNode(6)
	LevelOrderTraverse(node1)
}

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, left: nil, right: nil}
}

func PreOrderTraverse(node *TreeNode) {
	if nil == node {
		return
	}
	fmt.Println(node)
	PreOrderTraverse(node.left)
	PreOrderTraverse(node.right)
}

func InOrderTraverse(node *TreeNode) {
	if nil == node {
		return
	}
	InOrderTraverse(node.left)
	fmt.Println(node)
	InOrderTraverse(node.right)
}

func PostOrderTraverse(node *TreeNode) {
	if nil == node {
		return
	}
	PostOrderTraverse(node.left)
	fmt.Println(node)
	PostOrderTraverse(node.right)
}

func LevelOrderTraverse(node *TreeNode) {
	if nil == node {
		return
	}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		len1 := len(queue)
		for i := 0; i < len1; i++ {
			if len(queue) == 0 {
				continue
			}
			curr := queue[0]
			queue = queue[1:len(queue)]
			fmt.Print(curr.Val, "  ")
			if curr.left != nil {
				queue = append(queue, curr.left)
			}
			if curr.right != nil {
				queue = append(queue, curr.right)
			}
		}
	}

}
