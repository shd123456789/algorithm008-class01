func countSubstrings(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
        dp[i][i] = true 
    }
    re := n
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                if j - 1 == i {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i + 1][j - 1]
                }
            } else {
                dp[i][j] = false
            }
            if dp[i][j] {
                re++
            }
        }
    }
    return re
}