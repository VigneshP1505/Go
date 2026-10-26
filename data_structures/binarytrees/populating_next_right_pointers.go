package binarytrees

type TNode struct {
	Val       int
	Left      *TNode
	Right     *TNode
	NextRight *TNode
}

func PopulatingNextRightPointers(root *TNode) {
	if root == nil {
		return
	}
	queue := []TNode{*root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var prevNode *TNode
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			if prevNode != nil {
				prevNode.NextRight = &node
			}
			prevNode = &node

			if node.Left != nil {
				queue = append(queue, node)
			}

			if node.Right != nil {
				queue = append(queue, *node.Right)
			}
		}
		prevNode.NextRight = nil
	}
}

func LevelSize(root *TNode) {
	if root == nil {
		return
	}
	queue := []TNode{*root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for range levelSize {

		}
	}
}
