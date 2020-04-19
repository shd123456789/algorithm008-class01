// 26. 删除排序数组中的重复项

func removeDuplicates(nums []int) int {
    len := len(nums)
    if len == 0 {
        return 0
    }
    tmp := nums[len - 1] + 1
    step := 0
    for i := 0; i < len; i++ {
        if tmp == nums[i] { // 当相等时记录移动的总步数
            step++
        } else {
            nums[i - step] = nums[i]
            tmp = nums[i]
        }
    }
    return len - step
}
