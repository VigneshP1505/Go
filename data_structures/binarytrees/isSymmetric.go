package binarytrees

func isSymmetric(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	return node1.Val == node2.Val && isSymmetric(node1.Left, node2.Right) && isSymmetric(node1.Right, node2.Left)
}

func IsSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}
