// 时间复杂度O(n) 空间复杂度O(n) 3遍
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
    res = max(res, lh + rh)
    return max(lh, rh) + 1
}
func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}