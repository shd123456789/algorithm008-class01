func longestPalindrome(s string) string {
    n := len(s)
    if n < 2 {
        return s
    }
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
        dp[i][i] = true 
    }
    maxLen := 1
    start := 0
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
            if dp[i][j] && j - i + 1 >= maxLen {
                start = i
                maxLen = j - i + 1
            }
        }
    }
    return string(s[start:start+maxLen])
}