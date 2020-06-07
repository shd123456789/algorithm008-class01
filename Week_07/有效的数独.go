func isValidSudoku(board [][]byte) bool {
    col,row,box := make([][9]bool, 9),make([][9]bool, 9),make([][9]bool, 9)
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] != '.' {
                v := board[i][j] - '1'
                boxIdx := (i / 3) * 3 + j / 3
                if row[i][v] || col[j][v] || box[boxIdx][v] {
                    return false
                }
                row[i][v],col[j][v],box[boxIdx][v] = true, true, true
            }
        }
    }
    return true
}