var col,bie,na map[int]bool
var paths [][]int
var path []int
func solveNQueens(n int) [][]string {
    col,bie,na = make(map[int]bool),make(map[int]bool),make(map[int]bool)
    paths,path = [][]int{}, []int{}
    dfs(n , 0)
    return generateResult(n)
}
func dfs(n int, row int) {
   if row == n {
       paths = append(paths, append([]int{}, path...))
   }
   for i := 0; i < n; i++ {
        if col[i] || bie[row + i] || na[row - i] {
           continue
        }
        col[i], bie[row + i], na[row - i] = true,true,true
        path = append(path, i)
        dfs(n, row + 1)
        col[i], bie[row + i], na[row - i] = false,false,false
        path = path[:len(path) - 1]
   }
}
func generateResult(n int) [][]string {
    result := make([][]string, 0) 
    for _,v := range paths {
        s := make([]string, n)
        for k,v1 := range v {
            s1 := make([]byte, n)
            for i := 0; i < n; i++ {
                if i == v1 {
                    s1[i] = 'Q'
                } else {
                    s1[i] = '.'
                }
            }
            s[k] = string(s1)
        }
        result = append(result, s)
    }
    return result
}