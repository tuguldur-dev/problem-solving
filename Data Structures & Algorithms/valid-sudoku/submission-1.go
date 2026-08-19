func isValid(board [][]byte, _row int, _col int) bool {
	hashset := make(map[byte]bool)
	for row := _row; row < _row+3; row++ {
		for col := _col; col < _col+3; col++ {
			if board[row][col] == '.' {
				continue
			}
			if _, ok := hashset[board[row][col]]; ok {
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
			if !isValid(board, row, col) {
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
			if _, ok := hashset[board[row][col]]; ok {
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
			if _, ok := hashset[board[row][col]]; ok {
				return false
			}
			hashset[board[row][col]]=true
		}
	}
	return true
}
