/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    max := 1;
    for i := 0; i < len(root.Children); i++ {
        n := 1 + maxDepth(root.Children[i])
        if n > max {
            max = n
        }
    }
    return max
}