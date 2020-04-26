// 移动零
func moveZeroes(nums []int)  {
    insertPos := 0
    for i := 0; i < len(nums); i++ {
        if (nums[i] != 0) {
            tmp := nums[i]
            nums[i] = 0
            nums[insertPos] = tmp
            insertPos++
        }
    }
}