/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    return 1 + min(minDepth(root.Left), minDepth(root.Right))
}



func min(n1 int, n2 int) int {
    if n1 > n2 {
        return n2
    } else {
        return n1
    }
}