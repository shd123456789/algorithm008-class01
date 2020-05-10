// 3遍
func getKthMagicNumber(k int) int {
    p1,p2,p3 := 0,0,0
    dp := make([]int, k)
    dp[0] = 1
    for i := 1; i < k; i++ {
        dp[i] = min(dp[p1] * 3, dp[p2] * 5, dp[p3] * 7)
        if dp[i] == dp[p1] * 3 {
            p1++
        } 
        if dp[i] == dp[p2] * 5 {
            p2++
        } 
        if dp[i] == dp[p3] * 7 {
            p3++
        }
    } 
    return dp[k-1]
}

func min(n1 int,n2 int,n3 int) int {
    var min int
    if n1 > n2 {
        min = n2
    } else {
        min = n1
    }
    if min > n3 {
        min = n3
    }
    return min
}