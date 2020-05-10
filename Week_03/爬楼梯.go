// 3遍
// 动态规划 时间复杂度为O(n), 空间复杂度为O(1)
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    one,second := 1,2
    for i := 3; i <= n; i++ {
        one,second = second,second + one
    }
    return second
}

// 递归加备忘录方法 时间复杂度为O（N）空间复杂度为O（N）
var record []int 
func climbStairs(n int) int {
    record = make([]int, n)
    return climb(0, n)
}
func climb(s int, n int) int {
    if s > n {
        return 0
    }
    if s == n {
        return 1
    }
    if record[s] > 0 {
        return record[s]
    }
    record[s] = climb(s + 1, n) + climb(s + 2, n)
    return record[s]
}

