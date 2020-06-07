// 时间O(logn) 空间O(1)
func mySqrt(x int) int {
    res,left,right := 0,0,x
    for left <= right {
        mid := left + (right - left) >> 1
        if mid * mid <= x {
            res = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}