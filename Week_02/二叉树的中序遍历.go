/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    return helper(root, &res)
}

func helper(root *TreeNode, res *[]int) []int {
    if root != nil {
        helper(root.Left, res)
        *res = append(*res, root.Val)
        helper(root.Right, res)
    }
    return *res
}
// 遍历 2遍 //时间O(n)
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var res []int
    stack := []*TreeNode{root}
    for len(stack) != 0 {
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if node != nil {
            if node.Right != nil {
                stack = append(stack, node.Right)
            }
            stack = append(stack, node, nil) // nil作为表识符表示反问过了等后续加入结果集
            if node.Left != nil {
                stack = append(stack, node.Left)
            }
        } else {
            node = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            res = append(res, node.Val)
        }
    }
    return res
}




