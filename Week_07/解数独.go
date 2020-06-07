var col,row,box [][9]bool
func solveSudoku(board [][]byte)  {
    col,row,box = make([][9]bool, 9),make([][9]bool, 9),make([][9]bool, 9)
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] != '.' {
                v := board[i][j] - '1'
                boxIdx := (i / 3) * 3 + j / 3
                row[i][v],col[j][v],box[boxIdx][v] = true, true, true
            }
        }
    }
    dfs(board, 0, 0)
}

func dfs(board [][]byte, r int, c int) bool  {
    if r == 9 {        
        return true
    }
    newR,newC := r, c + 1
    if newC == 9 {
        newR,newC = r + 1,0
    }
    if board[r][c] == '.' {
        boxIdx := (r / 3) * 3 + c / 3
        for i := 1; i <= 9; i++ {
            v := i - 1
            if row[r][v] || col[c][v] || box[boxIdx][v] {
               continue
            } else {
                board[r][c] = byte(i) + '0'
                row[r][v],col[c][v],box[boxIdx][v] = true,true,true
                res := dfs(board, newR, newC)
                row[r][v],col[c][v],box[boxIdx][v] = false,false,false
                if res {
                    return true
                }
                board[r][c] = '.'
            }
        }
    } else {
        res := dfs(board, newR, newC)
        if res {
            return true
        }
    }
    return false
}