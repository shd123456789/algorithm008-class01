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

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil {
        return nil
    }

    type trackTreeNode struct{
        visited bool
        root *Node
    }
    var res []int
    var stack []*trackTreeNode
    stack = append(stack, &trackTreeNode{root: root})
    for len(stack) != 0 {
        node := stack[len(stack) - 1]
        if node.visited {
            res = append(res, node.root.Val)
            stack = stack[:len(stack) - 1]
            continue
        }
        for i := len(node.root.Children) - 1; i >= 0; i-- {
            stack = append(stack, &trackTreeNode{root: node.root.Children[i]})
        }
        node.visited = true
    }
    return res
}

