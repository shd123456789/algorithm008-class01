// 移动零 3遍 时间复杂度为O(n)空间复杂度为O(1)
func moveZeroes(nums []int)  {
    insertPos := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i],nums[insertPos] = 0,nums[i]
            insertPos++
        }
    }
}