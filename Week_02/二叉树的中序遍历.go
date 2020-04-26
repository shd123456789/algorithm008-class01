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
// 遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res []int
    var stack []*TreeNode

    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        index := len(stack) - 1
        res = append(res, stack[index].Val)
        root = stack[index].Right
        stack = stack[:index]
    }
    return res
}




