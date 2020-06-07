// 2遍 时间复杂度O（n） 空间O(1)
func sortColors(nums []int)  {
    p1,p2,curr := 0,len(nums) - 1,0
    for curr <= p2 {
        if nums[curr] == 0 {
            nums[p1],nums[curr] = nums[curr],nums[p1]
            p1++
            curr++
        } else if nums[curr] == 2 {
            nums[p2],nums[curr] = nums[curr],nums[p2]
            p2--
        } else {
            curr++
        }
    }
}