// 时间最好O1 最坏情况便利所有节点O(N), 空间最坏情况O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == p || root == q || root == nil {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left == nil {
        return right
    }
    if right == nil {
        return left
    }
    return root
}