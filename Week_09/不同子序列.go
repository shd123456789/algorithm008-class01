func numDistinct(s string, t string) int {
    dp := make([][]int, len(t) + 1)
    for i := range dp {
        dp[i] = make([]int, len(s) + 1)
    }
    for i := 0; i < len(dp[0]); i++ {
       dp[0][i] = 1
    }
    for i := 1; i < len(dp); i++ {
       dp[i][0] = 0
    }
    // t : i    s : j
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if s[j - 1] == t[i - 1] {
                dp[i][j] = dp[i - 1][j - 1] + dp[i][j - 1]
            } else {
                dp[i][j] = dp[i][j - 1]
            }
        } 
    }
    return dp[len(dp) - 1][len(dp[0]) - 1]
}