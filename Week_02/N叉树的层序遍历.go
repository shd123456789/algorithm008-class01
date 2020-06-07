/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
var res [][]int
func levelOrder(root *Node) [][]int {
    res = make([][]int, 0)
    return helper(root, 0)
}

func helper(root *Node, level int) [][]int{
    if root != nil {
        if len(res) > level  {
            res[level] = append(res[level], root.Val)
        } else {
            res = append(res, []int{root.Val})
        }
        for j := 0; j < len(root.Children); j++ {
            helper(root.Children[j], level + 1)
        }
    }
    return res
}
