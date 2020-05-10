// 2遍 时间复杂度为O(n),空间复杂度为系统栈消耗树的高度最好情况下是O(logn)最坏是O(n)
/*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left,root.Right = root.Right,root.Left
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}
