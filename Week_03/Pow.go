// 时间空间均为O(logn)
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    } else if n == 1 {
        return x
    } else if n < 0 {
        return 1 / myPow(x, -n)
    }
    half := myPow(x, n/2)
    if n % 2 == 0 {
        return half * half
    } else {
        return half * half * x
    }
}