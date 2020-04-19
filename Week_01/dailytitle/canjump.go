// 55. 跳跃游戏
func canJump(nums []int) bool { 
    len := len(nums)
    k := 0
    for i := 0; i < len; i++ {
        if k < i {
            return false
        }
        if k < (i + nums[i]) {
            k = i + nums[i]
        }
    }
    return true
}