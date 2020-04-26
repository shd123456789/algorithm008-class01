var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	helper(root)
	return res
}

func helper(root *TreeNode) int {
	if root == nil {
        return 0
    }
    lh := helper(root.Left)
    rh := helper(root.Right)
    res = max(lh + rh, res)
    return max(lh, rh) + 1
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}