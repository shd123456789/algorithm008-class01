/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    return helper(root, &res)
}

func helper(root *TreeNode, res *[]int) []int {
    if root != nil {
        *res = append(*res, root.Val)
        helper(root.Left, res)
        helper(root.Right, res)
    }
    return *res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
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
            if node.Left != nil {
                stack = append(stack, node.Left)
            }
            stack = append(stack, node, nil) // nil作为表识符表示反问过了等后续加入结果集
        } else {
            node = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            res = append(res, node.Val)
        }
    }
    return res
}
