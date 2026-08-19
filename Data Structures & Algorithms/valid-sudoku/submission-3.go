func isDuplicated(hashmap map[byte]bool, board [][]byte, row int, col int) bool {
    if _, ok := hashmap[board[row][col]]; ok {
        return true
    }
    return false
}

func isValidBox(board [][]byte, _row int, _col int) bool {
	hashset := make(map[byte]bool)
	for row := _row; row < _row+3; row++ {
		for col := _col; col < _col+3; col++ {
			if board[row][col] == '.' {
				continue
			}
			if isDuplicated(hashset, board, row, col) {
				return false
			}
			hashset[board[row][col]]=true
		}
	} 
	return true
}
func isValidSudoku(board [][]byte) bool {
	n := len(board)
	for row := 0; row < n; row+=3 {
		for col := 0; col < n; col+=3 {
			if !isValidBox(board, row, col) {
				return false
			}
		}
	}
	for row := 0; row < n; row++ {
		hashset := make(map[byte]bool)
		for col := 0; col < n; col++ {
			if board[row][col] == '.' {
				continue
			}
			if isDuplicated(hashset, board, row, col) {
				return false
			}
			hashset[board[row][col]]=true
		}
	}
	for col := 0; col < n; col++ {
		hashset := make(map[byte]bool)
		for row := 0; row < n; row++ {
			if board[row][col] == '.' {
				continue
			}
			if isDuplicated(hashset, board, row, col) {
				return false
			}
			hashset[board[row][col]]=true
		}
	}
	return true
}
