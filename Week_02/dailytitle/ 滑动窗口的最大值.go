func maxSlidingWindow(nums []int, k int) []int {
    var queue,res []int
    for i := 0; i < len(nums); i++ {
        for len(queue) > 0 && nums[i] >= nums[queue[len(queue) - 1]] { // 保持单调队列递减
            queue = queue[:len(queue) - 1]
        }
        queue = append(queue, i)
        if queue[0] < i + 1 - k {
            queue = queue[1:len(queue)]
        }
        if i+1 >= k {
            res = append(res, nums[queue[0]])
        } 
    }
    return res
}
