/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return helper(root) != -1
}

func helper(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHight := helper(root.Left)
    if leftHight == -1 {
        return -1
    }
    rightHight := helper(root.Right)
    if rightHight == -1 {
        return -1
    }
    if abs(leftHight - rightHight) >= 2 {
        return - 1
    }
    return 1 + max(leftHight, rightHight)
}

func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}

func abs(n int) int {
    if n > 0 {
        return n
    } else {
        return -n
    }
}