func numIslands(grid [][]byte) int {
    sum := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                sum ++
                dfs(i, j, grid)
            }
        }
    }
    return sum
}

func dfs(i int, j int, grid[][]byte) {
    grid[i][j] = '0'
    if i - 1 >= 0 && grid[i - 1][j] == '1' {
        dfs(i - 1, j, grid)
    }
    if i + 1 < len(grid) && grid[i + 1][j] == '1' {
        dfs(i + 1, j, grid)
    }
    if j - 1 >= 0 && grid[i][j - 1] == '1' {
        dfs(i, j - 1, grid)
    }
    if j + 1 < len(grid[0]) && grid[i][j + 1] == '1' {
        dfs(i, j + 1, grid)
    }
}


