package binarytrees

// construct uniq

// construct sorted BST recursively
func constructBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = constructBST(nums[:mid])
	root.Right = constructBST(nums[mid:])
	return root
}

// construct unsorted BST recursively
func constructBSTUnsorted(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = constructBSTUnsorted(root.Left, val)
	} else {
		root.Right = constructBSTUnsorted(root.Right, val)
	}
	return root
}

func buildBSTUnsorted(nums []int) {
	var root *TreeNode
	for i := 0; i < len(nums); i++ {
		root = constructBSTUnsorted(root, nums[i])
	}
}
