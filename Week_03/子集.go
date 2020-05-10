// 时间复杂O(2^n),空间复杂度O(n)
var res [][]int
var path []int
func subsets(nums []int) [][]int {
    res,path = make([][]int, 0), make([]int, 0)
    backtrack(0, nums)
    return res
}

func backtrack(level int, nums []int) {
    if level == len(nums) {
        res = append(res, append([]int{}, path...))
        return
    }

    backtrack(level + 1, nums)
    path = append(path, nums[level])
    backtrack(level + 1, nums)
    path = path[:len(path) - 1]
}