var res []string
func generateParenthesis(n int) []string {
    res = make([]string, 0)
    dfs(0, 0, n, "")
    return res
}

func dfs(left int, right int, n int, str string) {
    if 2 * n == len(str) {
        res = append(res, str)
    }
    if left < n {
        dfs(left + 1, right, n, str + "(")
    }
    if right < left {
        dfs(left, right + 1, n, str + ")")
    }
}