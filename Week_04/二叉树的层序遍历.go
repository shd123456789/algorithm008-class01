// 时间复杂度为O(n) 空间复杂度O(n)
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        tmp := []int{}
        l := len(queue)
        for i := 0; i < l; i++ {
            node := queue[i]
            tmp = append(tmp, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[l:]
        res = append(res, tmp)
    }
    return res
}