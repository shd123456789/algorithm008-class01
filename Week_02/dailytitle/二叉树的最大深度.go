// 时间O(n) 空间最坏O(n) 3遍
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(n1 int, n2 int) int {
    if n1 < n2 {
        return n2
    } else {
        return n1
    }
}