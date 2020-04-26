/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    res := make([]int, 0)
    return helper(root, &res)
}

func helper(root *Node, res *[]int) []int {
    if root != nil {
        *res = append(*res, root.Val)
        for i := 0; i < len(root.Children); i++ {
            helper(root.Children[i], res)
        }
    }
    return *res
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return nil
    }
    var res []int
    var stack []*Node
    stack = append(stack, root)
    for len(stack) != 0 {
        node := stack[len(stack) - 1] 
        res = append(res, node.Val)
        stack = stack[:len(stack) - 1] 
        for i := len(node.Children) - 1; i >= 0; i-- {
            stack = append(stack, node.Children[i])
        }
    } 
    return res
}

