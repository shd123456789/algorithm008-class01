// 数组加1
func plusOne(digits []int) []int { 
    len := len(digits) - 1
    for i := len; i >=0; i-- {
        v := digits[i] + 1
        digits[i] = v % 10
        if digits[i] != 0 {
            return digits
        }
    }
    return append([]int{1}, digits...) 
}