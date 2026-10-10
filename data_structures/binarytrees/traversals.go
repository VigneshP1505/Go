package binarytrees

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Val)
}

func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)
}
