// dfs 深度优先
func numIslands(grid [][]byte) int {
    nr := len(grid)
    if nr == 0 {
        return 0
    }
    nc := len(grid[0])
    sumLands := 0
    for i := 0; i < nr; i++ {
        for j := 0; j < nc; j++ { 
            if grid[i][j] == '1' {
                sumLands++
                dfs(grid, i, j)
            }
        }
    }
    return sumLands
}

func dfs(grid [][]byte,r int,c int) {
    grid[r][c] = '0'
    nr,nc := len(grid),len(grid[0])

    if r + 1 < nr && grid[r + 1][c] == '1' {
        dfs(grid, r + 1, c)
    }
    if c + 1 < nc && grid[r][c + 1] == '1' {
        dfs(grid, r, c + 1)
    }
    if r - 1 >= 0 && grid[r - 1][c] == '1'{
        dfs(grid, r - 1, c)
    }
    if c - 1 >= 0 && grid[r][c - 1] == '1' {
        dfs(grid, r, c - 1)
    }
}