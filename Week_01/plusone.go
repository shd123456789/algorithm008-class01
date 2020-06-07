// 数组加1 时间空间最好情况为O（1）最坏为O(n) 3遍
func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i] = digits[i] + 1
        if digits[i] == 10 {
            digits[i] = 0
        } else {
            return digits
        }
    }
    return append([]int{1}, digits...)
}
   
