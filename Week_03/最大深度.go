
// 2遍 时间复杂度为O(n),空间复杂度为树的高度最好情况下是O(logn)最坏是O(n)
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}