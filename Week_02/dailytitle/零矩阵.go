func setZeroes(matrix [][]int)  {
    row := make(map[int]byte)
    col := make(map[int]byte)
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                row[i] = 1
                col[j] = 1
            }
        }
    }
    for i,_ := range row {
        for j := 0; j < len(matrix[0]); j++ {
            matrix[i][j] = 0
        }
    }
    for j,_ := range col {
        for i := 0; i < len(matrix); i++ {
            matrix[i][j] = 0
        }
    }
}