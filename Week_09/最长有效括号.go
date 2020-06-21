func longestValidParentheses(s string) int {
    if s == "" {
        return 0
    }
    dp := make([]int, len(s))
    dp[0] = 0
    maxL := 0
    for i := 1; i < len(s); i++ {
        if s[i] == ')' {
            if s[i - 1] == '(' {
                dp[i] = 2
                if i - 2 >= 0 {
                    dp[i] += dp[i - 2]
                }   
            } else if (i - dp[i - 1]) >= 1 && s[i - dp[i - 1] - 1] == '(' {
                dp[i] = dp[i-1] + 2
                if i - dp[i - 1] >= 2 {
                    dp[i] += dp[i - dp[i - 1] - 2]
                }
            }
        }
        maxL = max(maxL, dp[i])
    }
    return maxL
}


//  s[i] == '(' dp[i] = 0   
//  s[i] == ')'  if s[i - 1] == '('  dp[i] = dp[i - 2] + 2
//  s[i] == ')'  if s[i - 1] == ')' && s[i - dp[i - 1] - 1] == '('
// dp[i-1] + 2 + dp[i - dp[i - 1] - 2]

func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}