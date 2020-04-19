// 70. 爬楼梯 动态规划
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    first := 1
    second := 2
    for i := 3; i <= n; i++ {
        method := first + second
        first = second
        second = method
    }
    return second
}
