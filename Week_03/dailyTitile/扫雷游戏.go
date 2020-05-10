func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	helper(board, click)
	return board
}
func helper(board [][]byte, click []int) {
	if !isBorderRight(board, click) || board[click[0]][click[1]] != 'E' {
		return
	} 
	var count byte
    aroundPos := getAroundPos(click)
    for _,v := range aroundPos { // 获取周围的炸弹数
        count += bombNum(board, v)
    }
	if count > 0 {
		board[click[0]][click[1]] = count + '0'
	} else {
		board[click[0]][click[1]] = 'B'
        for _,v := range aroundPos { // 由于是B扩展周围
            helper(board, v)
        }
	}
}

func getAroundPos (click []int) [][]int{
    return [][]int{
        []int{click[0] + 1, click[1]}, []int{click[0] - 1, click[1]}, 
        []int{click[0], click[1] + 1}, []int{click[0], click[1] - 1},
        []int{click[0] + 1, click[1] + 1}, []int{click[0] + 1, click[1] - 1},
        []int{click[0] - 1, click[1] - 1}, []int{click[0] - 1, click[1] + 1},
    }
}
func bombNum(board [][]byte, click []int) byte {
    var count byte
	if !isBorderRight(board, click) {
		count = byte(0)
	} else if board[click[0]][click[1]] == 'M' {
		count = byte(1)
	}
    return count
}
func isBorderRight(board [][]byte, click []int) bool {
    if click[0] >= len(board) || click[0] < 0 || click[1] >= len(board[0]) || click[1] < 0 {
        return false
    }
    return true
}