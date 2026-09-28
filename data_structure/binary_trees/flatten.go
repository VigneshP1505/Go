package binarytrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	curr := root
	for curr != nil {
		if curr.Left != nil {
			predecessor := curr.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
	}
	curr = curr.Right
}

type Nodes struct {
	Val   int
	Left  *Nodes
	Right *Nodes
}

func _flatten(root *Nodes) *Nodes {
	if root == nil {
		return root
	}

}
