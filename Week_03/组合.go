// 2遍
var res [][]int
var element []int
func combine(n int, k int) [][]int {
	res = make([][]int, 0)
    element = make([]int, 0, k)
    backtrack(1, n, k)
    return res
}

func backtrack(p int, n int, k int) {
    if len(element) == k {
        res = append(res, append([]int{}, element...))
    }

    for i := p; i <= n; i++ {
        if k - len(element) > n - i + 1 { // 剪枝减少递归次数意思是当i遍历到一定数值时，由于在继续遍历下去也无法凑到k个数
            return 
        }
        element = append(element, i) // 处理当前层
        backtrack(i + 1, n, k) // 下探到下一层
        element = element[:len(element) - 1] // 状态清除
    }
}