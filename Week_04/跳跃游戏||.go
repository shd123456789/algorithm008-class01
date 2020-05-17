// 时间复杂度为O(n) 空间复杂度为O(1)
func jump(nums []int) int {
    res,end,maxPos := 0,0,0
    for i := 0; i < len(nums) - 1; i++ {
        maxPos = max(maxPos, nums[i] + i)
        if end == i {
            end = maxPos
            res++
        }
    }
    return res
}

func max (n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}