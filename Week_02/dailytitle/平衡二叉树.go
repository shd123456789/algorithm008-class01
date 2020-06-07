// 时间O(n) 空间O(logn)
func isBalanced(root *TreeNode) bool {
    return helper(root) != -1
}

func helper(root *TreeNode) int {
    if root == nil {
        return 0
    }
    lh := helper(root.Left)
    if lh == -1 {
        return -1
    }
    rh := helper(root.Right)
    if rh == -1 {
        return -1
    }
    if abs(lh - rh) >= 2 {
        return -1
    }
    return 1 + max(lh, rh)
}

func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}

func abs(n int) int {
    if n > 0 {
        return n
    } else {
        return -n
    }
}

【day38】

https://leetcode-cn.com/problems/word-ladder/description/


