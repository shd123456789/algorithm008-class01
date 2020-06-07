// 时间复杂度O(n) 空间复杂度为O(1) 2遍
func addDigits(num int) int {
    res := 0
    for num != 0 {
        res += num % 10
        if num = num / 10; num == 0 {
            if res < 10 {
                break
            }
            num,res = res,0
        }
    }
    return res
}

// O(1)
func addDigits(num int) int {
    return (num - 1) % 9 + 1
}   