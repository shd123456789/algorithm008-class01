// 2遍  时间复杂度为O(n),空间复杂度为系统栈消耗树的高度最好情况下是O(logn)最坏是O(n)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}

func helper (root *TreeNode, low int, high int) bool {
    if root == nil {
        return true
    }
    if root.Val <= low || root.Val >= high {
        return false
    }
    return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
}