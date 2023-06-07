package main


func removeLAndT(board [][]int) int {
	// 首先处理水平方向上的 "L" 和 "T"
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 0 {
				// 判断 "L" 形式
				if i+2 < len(board) && j+1 < len(board[0]) &&
					board[i][j] == board[i+1][j] &&
					board[i+1][j] == board[i+2][j] &&
					board[i+1][j] == board[i+1][j+1] {
					board[i][j] = 0
					board[i+1][j] = 0
					board[i+2][j] = 0
					board[i+1][j+1] = 0
					return 1 + removeLAndT(board)
				}
				// 判断 "T" 形式
				if i+2 < len(board) && j-1 >= 0 &&
					board[i][j] == board[i+1][j] &&
					board[i+1][j] == board[i+2][j] &&
					board[i+1][j] == board[i+1][j-1] {
					board[i][j] = 0
					board[i+1][j] = 0
					board[i+2][j] = 0
					board[i+1][j-1] = 0
					return 1 + removeLAndT(board)
				}
			}
		}
	}
	// 处理垂直方向上的 "L" 和 "T"
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 0 {
				// 判断 "L" 形式
				if i+1 < len(board) && j+2 < len(board[0]) &&
					board[i][j] == board[i][j+1] &&
					board[i][j+1] == board[i][j+2] &&
					board[i][j+1] == board[i+1][j+1] {
					board[i][j] = 0
					board[i][j+1] = 0
					board[i][j+2] = 0
					board[i+1][j+1] = 0
					return 1 + removeLAndT(board)
				}
				// 判断 "T" 形式
				if i-1 >= 0 && j+2 < len(board[0]) &&
					board[i][j] == board[i-1][j+1] &&
					board[i-1][j+1] == board[i][j+1] &&
					board[i][j+1] == board[i][j+2] {
					board[i][j] = 0
					board[i-1][j+1] = 0
					board[i][j+1] = 0
					board[i][j+2] = 0
					return 1 + removeLAndT(board)
				}
			}
		}
	}
	return 0
}

