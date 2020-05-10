// 时间复杂度为O(n) 空间复杂度为O(n)
 var index int
 var hashPos map[int]int
func buildTree(preorder []int, inorder []int) *TreeNode {
    hashPos = make(map[int]int)
    index = 0
    for k,v := range inorder {
        hashPos[v] = k
    }
    return helper(0, len(inorder), preorder)
}

func helper(left int, right int, preorder []int) *TreeNode {
    if left == right {
        return nil
    }
    val := preorder[index]
    pos := hashPos[val]
    index++
    return &TreeNode{
        Val : val,
        Left : helper(left, pos, preorder),
        Right : helper(pos + 1, right, preorder),
    }
}