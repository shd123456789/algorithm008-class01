// 时间O(logn) 空间O(1)
func isPerfectSquare(num int) bool {
    l,r,mid := 0,num,0
    for l <= r {
        mid = l + (r - l) >> 2
        sqrt := mid * mid
        if sqrt == num {
            return true
        } else if sqrt < num {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return false
}