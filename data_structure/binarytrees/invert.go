package binarytrees

type InvertNode struct {
	Val   int
	Left  *InvertNode
	Right *InvertNode
}

func invertUtil(node *InvertNode) *InvertNode {
	if node == nil {
		return nil
	}
	left := invertUtil(node.Left)
	right := invertUtil(node.Right)
	temp := left
	node.Left = right
	node.Right = temp
	return node
}

func Invert(root *InvertNode) {
	invertUtil(root)
}
