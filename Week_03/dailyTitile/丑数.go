func nthUglyNumber(n int) int {
    p1,p2,p3 := 0,0,0
    dp := make([]int, n)
    dp[0] = 1 
    for i := 1; i < n; i++ {
        n1,n2,n3 := dp[p1]*2, dp[p2]*3, dp[p3]*5
        dp[i] = min(n1, n2, n3)
        if dp[i] == n1 {
            p1++
        }
        if dp[i] == n2 {
            p2++
        }
        if dp[i] == n3 {
            p3++
        }
    }
    return dp[n - 1]
}

func min(n1 int, n2 int, n3 int) int {
    if n1 > n2 {
        n1 = n2
    }
    if n1 > n3 {
        n1 = n3
    }
    return n1
}