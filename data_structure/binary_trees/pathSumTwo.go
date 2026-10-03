package binarytrees

type PathSumIINode struct {
	Val   int
	Left  *PathSumIINode
	Right *PathSumIINode
}

func pathSumIIUtil(root *PathSumIINode, targetSum int, path []PathSumIINode, paths *[][]PathSumIINode) {
	if root == nil {
		return
	}
	path = append(path, *root)
	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			currentPath := make([]PathSumIINode, len(path))
			copy(currentPath, path)
			*paths = append(*paths, currentPath)
		}
		return
	}
	pathSumIIUtil(root.Left, targetSum, path, paths)
	pathSumIIUtil(root.Right, targetSum, path, paths)
}

func PathSumII(root *PathSumIINode, targetSum int) [][]PathSumIINode {
	var paths [][]PathSumIINode
	pathSumIIUtil(root, targetSum, nil, &paths)
	return paths
}
