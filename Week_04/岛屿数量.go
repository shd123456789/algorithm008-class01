// 时间复杂度为O(MN) 空间复杂度为O(MN) 3遍
func numIslands(grid [][]byte) int {
    res := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                res++
                dfs(grid, i, j)
            }
        }
    }
    return res
}

func dfs(grid [][]byte,r int,c int) {
    grid[r][c] = '0'
    l1,l2 := len(grid),len(grid[0])
    if r - 1 >= 0 && grid[r-1][c] == '1' {
        dfs(grid, r-1, c)
    }
    if r + 1 < l1 && grid[r + 1][c] == '1' {
        dfs(grid, r+1, c)
    }
    if c - 1 >= 0 && grid[r][c - 1] == '1' {
        dfs(grid, r, c - 1)
    }
    if c + 1 < l2 && grid[r][c + 1] == '1' {
        dfs(grid, r, c + 1)
    }
}