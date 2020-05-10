var res [][]int
var used []int
var path []int
func permuteUnique(nums []int) [][]int {
    res,used,path = make([][]int, 0), make([]int, len(nums)), make([]int, 0)
    dfs(nums)
    return res
}

func dfs(nums []int) {
    if len(path) == len(nums) {
       res = append(res, append([]int{}, path...))
    }
    record := make(map[int]bool)
    for i,v := range nums {
        if used[i] == 1 || record[v] {
           continue
        }
        record[v] = true
        used[i] = 1
        path = append(path, v)
        dfs(nums)
        used[i] = 0
        path = path[:len(path) - 1]
   }
}