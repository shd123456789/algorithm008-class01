/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    res := make([]int, 0)
    return helper(root, &res)
}

func helper(root *Node, res *[]int) []int {
    if root != nil {
        for i := 0; i < len(root.Children); i++ {
            helper(root.Children[i], res)
        }
        *res = append(*res, root.Val)
    }
    return *res
}

// 栈的实现
func postorder(root *Node) []int {
    if root == nil {
        return nil
    }
    var res []int
    stack := []*Node{root}
    for len(stack) != 0 {
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if node != nil {
            stack = append(stack, node, nil) // nil作为表识符表示反问过了等后续加入结果集
            for i := len(node.Children) - 1; i >= 0; i-- {
                stack = append(stack, node.Children[i])
            }
        } else {
            node = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            res = append(res, node.Val)
        }
    }
    return res
}

