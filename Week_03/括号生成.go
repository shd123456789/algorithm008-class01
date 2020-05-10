// 2遍
func generateParenthesis(n int) []string {
    res := make([]string, 0)
    dfs(0, 0, n, "", &res)
    return res
}

func dfs(left int, right int, n int, str string, res *[]string) {
    if 2 * n == len(str) {
        *res = append(*res, str)
        return 
    }
    if left < n {
        dfs(left + 1, right, n, str + "(", res)
    }
    if right < left {
        dfs(left, right + 1, n, str + ")", res)
    }
}