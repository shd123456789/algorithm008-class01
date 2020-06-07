// 时间复杂度O(nlogn) 空间复杂度结果集不算O(1)
func minimumAbsDifference(arr []int) [][]int {
    var res [][]int
    if len(arr) < 2 {
        return res
    }
    sort.Ints(arr)
    minDiff := arr[1] - arr[0] 
    for i := 1; i < len(arr) - 1; i++ { // 找最小差值
        if diff := arr[i + 1] - arr[i]; diff < minDiff {
            minDiff = diff
        }
    }
    for i := 0; i < len(arr) - 1; i++ {
        if arr[i + 1] - arr[i] == minDiff {
            res = append(res, []int{arr[i], arr[i + 1]})
        }
    }
    return res
}