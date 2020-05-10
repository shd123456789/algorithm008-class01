// 时间复杂度为O(n) 空间复杂度为O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    val := postorder[len(postorder) - 1]
    inPos := 0
    for k,v := range inorder {
        if v == val {
            inPos = k
            break
        }
    }
    rl := len(inorder) - inPos - 1 // 右子树长度
    return &TreeNode{
        Val : val,
        Left : buildTree(
            inorder[:inPos], 
            postorder[:len(postorder) - rl - 1],
        ),
        Right : buildTree(
            inorder[inPos + 1:], 
            postorder[len(postorder) - rl - 1:len(postorder) - 1],
        ),
    }
}