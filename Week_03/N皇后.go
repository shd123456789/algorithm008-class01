var col,bie,na map[int]bool
var path []int
var res [][]int
func solveNQueens(n int) [][]string {
    path,res = make([]int, 0),make([][]int, 0)
    col,bie,na = make(map[int]bool),make(map[int]bool),make(map[int]bool)
    dfs(n, 0)
    return generateResult(n)
}

func dfs(n int, row int) {
    if row == n {
        res = append(res, append([]int{}, path...))
        return 
    }
    for i := 0; i < n; i++ {
        if col[i] || bie[row + i] || na[row - i] {
            continue
        }
        col[i],bie[row + i],na[row - i] = true,true,true
        path = append(path, i)
        dfs(n, row + 1)
        col[i],bie[row + i],na[row - i] = false,false,false
        path = path[:len(path) - 1]
    }
}
func generateResult(n int) [][]string {
    result := make([][]string, len(res)) 
    for k,v := range res {
        s := make([]string, n)
        for k1,v1 := range v {
            s1 := make([]byte, n)
            for i := 0; i < n; i++ {
                if i == v1 {
                    s1[i] = 'Q'
                } else {
                    s1[i] = '.'
                }
            }
            s[k1] = string(s1)
        }
        result[k] = s
    }
    return result
}
