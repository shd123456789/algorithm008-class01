func isMatch(s string, p string) bool {
    dp := make([][]bool, len(s) + 1)
    for i := range dp {
        dp[i] = make([]bool, len(p) + 1)
    }
    dp[0][0] = true 
    for i := 1; i < len(dp[0]); i++ {
        dp[0][i] =  dp[0][i - 1] && p[i - 1] == '*'
    }
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if s[i - 1] == p[j - 1] || p[j - 1] == '?' {
                dp[i][j] = dp[i - 1][j - 1]
            } else if p[j - 1] == '*' {
                // adceb  逆思维推倒当e和*匹配时*可以和空匹配
                // *a*b   所以保证“adce”和“*a”为true匹配即可 dp[i][j - 1] 跳过"*"
                //        *可以和e匹配所以保证adc和*a*为true即可 dp[i - 1][j] 将“*”保留和前面字符匹配
                dp[i][j] = dp[i - 1][j] || dp[i][j - 1]
            }
        }
    }
    return dp[len(dp) - 1][len(dp[0]) - 1]
}