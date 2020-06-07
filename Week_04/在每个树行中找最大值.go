// 时间O(n） 空间O(n)
func largestValues(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        l := len(queue)
        max := queue[0].Val
        for i := 0; i < l; i++ {
            node := queue[i]
            if max < queue[i].Val {
                max = queue[i].Val
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[l:]
        res = append(res, max)
    }
    return res
}