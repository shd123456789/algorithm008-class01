// 时间复杂度O(MN)空间复杂O(MN) 2遍
func updateBoard(board [][]byte, click []int) [][]byte {
    if board[click[0]][click[1]] == 'M' {
        board[click[0]][click[1]] = 'X'
        return board
    }
    dfs(board, click)
    return board
}

func dfs(board [][]byte, click []int) {
    clickR,clickC := click[0],click[1]
    if isOverLimit(board, click) || board[clickR][clickC] != 'E' {
        return
    }
    aroundPos := getAroundPos(click)
    count := byte(0)
    for _,pos := range aroundPos {
        count += getBoomNums(board, pos)
    }
    if count > byte(0) {
        board[clickR][clickC] = count + '0'
    } else {
        board[clickR][clickC] = 'B'
        for _,pos := range aroundPos {
            dfs(board, pos)
        }
    }
}

func getAroundPos(click []int) [][]int {
    clickR,clickC := click[0],click[1]
    return [][]int{
        {clickR - 1, clickC},{clickR + 1, clickC},
        {clickR, clickC - 1},{clickR, clickC + 1},
        {clickR - 1, clickC - 1},{clickR - 1, clickC + 1},
        {clickR + 1, clickC - 1},{clickR + 1, clickC + 1},
    }
}

func getBoomNums(board [][]byte, click []int) byte {
    clickR,clickC := click[0],click[1]
    if !isOverLimit(board, click) && board[clickR][clickC] == 'M' { 
        return byte(1)
    }
    return byte(0)
}

func isOverLimit(board [][]byte, click []int) bool {
    l1,l2 := len(board) - 1, len(board[0]) - 1
    clickR,clickC := click[0],click[1]
    if clickR < 0 || clickR > l1 || clickC < 0 || clickC > l2 {
        return true
    }
    return false
}