// 时间复杂度O(n) 空间复杂度最坏O(n) 2遍
func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    maxDeep := 0
    for i := 0; i < len(root.Children); i++ {
        maxDeep = max(maxDeep, maxDepth(root.Children[i]))
    }
    return maxDeep + 1
}

func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}
