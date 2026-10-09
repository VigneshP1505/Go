package binarytrees

func ZigzagTraversal(root *Node) {
	if root == nil {
		return
	}

	paths := [][]int{}
	queue := []Node{*root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}
		}
		leftToRight = !leftToRight
		paths = append(paths, level)
	}
}

func bfs(root *Node) {
	if root == nil {
		return
	}
	queue := []Node{*root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, *node.Left)
		}
		if node.Right != nil {
			queue = append(queue, *node.Right)
		}
	}
}
