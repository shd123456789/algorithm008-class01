func nthUglyNumber(n int) int {
    p1,p2,p3 := 0,0,0
    dp := make([]int, n)
    dp[0] = 1
    for i := 1; i < n; i++ {
        n1,n2,n3 := dp[p1] * 2, dp[p2] * 3, dp[p3] * 5
        dp[i] = min(n1, n2, n3)
        if n1 == dp[i] {
            p1++
        }
        if n2 == dp[i] {
            p2++
        }
        if n3 == dp[i] {
            p3++
        }
    }
    return dp[n - 1]
}

func min(n1 int, n2 int, n3 int) int {
    n := 0
    if n1 > n2 {
        n = n2
    } else {
        n = n1
    }
    if n > n3 {
        n = n3
    }
    return n 
}