/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    return helper(root, &res)
}

func helper(root *TreeNode, res *[]int) []int {
    if root != nil {
        helper(root.Left, res)
        helper(root.Right, res)
        *res = append(*res, root.Val)
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
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    type trackTreeNode struct {
        visited bool
        root *TreeNode
    }

    var stack []*trackTreeNode
    var res []int
    stack = append(stack, &trackTreeNode{root: root})
    for len(stack) != 0 {
        node := stack[len(stack) - 1]
        if node.visited {
            res = append(res, node.root.Val)
            stack = stack[:len(stack) - 1]
            continue
        }
        if node.root.Right != nil {
            stack = append(stack, &trackTreeNode{root: node.root.Right})
        }
        if node.root.Left != nil {
            stack = append(stack, &trackTreeNode{root: node.root.Left})
        }
        node.visited = true
    }
    return res
}