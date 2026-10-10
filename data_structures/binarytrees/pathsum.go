package binarytrees

type PathSumNode struct {
	Val   int
	Left  *PathSumNode
	Right *PathSumNode
}

func pathSumUtil(root *PathSumNode, target int, sum int) bool {
	if root == nil {
		return false
	}
	sum = root.Val
	if root.Left == nil && root.Right == nil {
		return sum == target
	}
	return pathSumUtil(root.Left, target, sum) || pathSumUtil(root.Right, target, sum)
}

func PathSum(root *PathSumNode) bool {
	return pathSumUtil(root, 4, 0)
}
