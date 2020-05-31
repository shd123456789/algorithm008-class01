func minPathSum(grid [][]int) int {
    for i := 1; i < len(grid[0]); i++ { // 初始化第一行
        grid[0][i] = grid[0][i] + grid[0][i - 1]
    }

    for i := 1; i < len(grid); i++ {    // 初始化第一列
        grid[i][0] = grid[i][0] + grid[i - 1][0]
    }
    
    for i := 1; i < len(grid); i++ {
        for j := 1; j < len(grid[0]); j++ {
            grid[i][j] = grid[i][j] + min(grid[i - 1][j], grid[i][j - 1])
        }
    }
    return grid[len(grid) - 1][len(grid[0]) - 1]
}

func min(n1 int, n2 int) int {
    if n1 > n2 {
        return n2
    } else {
        return n1
    }
}