func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int, len(nums))
    dp[0] = 1
    maxL := 1
    for i := 1; i < len(nums); i++ {
        for j := i - 1; j >= 0; j-- {
            if nums[j] < nums[i] {
                dp[i] = max(dp[j], dp[i])
            }
        }
        dp[i] += 1
        maxL = max(dp[i], maxL)
    }
    return maxL
}

func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}