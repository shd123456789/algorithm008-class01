var count int
func totalNQueens(n int) int {
	count = 0
    dfs(n, 0, 0, 0, 0)
    return count
}

func dfs(n, row, col, pie, na int){
	if n == row {
        count++
        return
    }
    bits := (^(col | pie | na)) & ((1 << n) - 1) 
    for bits > 0 {
        pick := bits & -bits
        dfs(n, row + 1, col | pick, (pie | pick) << 1, (na | pick) >> 1)
        bits = bits & (bits - 1)
    }
}