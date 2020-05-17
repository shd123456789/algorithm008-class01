// 空间复杂度为O(1) 时间复杂度为O(n) 2遍
func canJump(nums []int) bool { 
    maxPos := 0
    for i := 0; i < len(nums); i++ {
        if maxPos < i {
            return false
        }
        if maxPos < i + nums[i] {
            maxPos = i + nums[i]
        }
    }
    return true
}
