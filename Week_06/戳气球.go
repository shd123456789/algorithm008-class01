func maxCoins(nums []int) int {
    n := len(nums)
    nums = append(nums, 1)
    nums = append([]int{1}, nums...)

    dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

    for i := n; i >= 0; i-- {
        for j := i + 1; j < n + 2; j++ {
            for k := i + 1; k < j; k++ {
                dp[i][j] = max(
                    dp[i][j], 
                    dp[i][k] + dp[k][j] + nums[i] * nums[j] * nums[k]);
            }
        }
    }
    return dp[0][n + 1];
}





func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}