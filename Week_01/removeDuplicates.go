// 26. 删除排序数组中的重复项(2遍) 时间复杂度为O(n),空间复杂度为O(1)
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    step := 0
    for i := 0; i < len(nums); i++ {
        if nums[step] != nums[i] {
            step++
            nums[step] = nums[i]
        }
    }
    return step + 1
}
    