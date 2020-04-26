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
    var res []int
    var stack []*TreeNode
    for len(stack) > 0 || root != nil {
        for root != nil {
            res = append(res, root.Val)
            stack = append(stack, root.Right)
            root = root.Left
        }
        index := len(stack) - 1
        root = stack[index]
        stack = stack[:index]
    }
    return res
}
