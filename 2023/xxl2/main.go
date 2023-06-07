package main

import (
	"fmt"
	"math/rand"
	"time"
)


/*该算法使用双重循环遍历棋盘中的每个格子，对于非空格子，分别进行水平和垂直方向的判断，判断是否可以进行连续消除。
如果可以进行连续消除，就将对应的格子标记为空，然后进行填充空缺的操作，将上面的格子往下移动，直到所有空缺都被填满为止。
如果没有可以进行连续消除的格子，算法结束。

在本算法中，我们使用 `board` 数组来表示棋盘，每个元素的值表示该格子中的颜色，其中 0 表示该格子为空。
在每次消除后，我们将对应格子的值设置为 0，然后再进行空缺的填充。我们还定义了一个 `randomColor` 函数，用于随机生成每个格子的颜色。

该算法实现了水平和垂直方向的连续消除，对于 L 形和 T 形的连续消除暂不支持，但可以通过进一步修改算法来实现。
*/

const (
	boardSize = 7
)

var (
	board [boardSize][boardSize]int
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initializeBoard()
	//printBoard()

	for eliminate() {
		printBoard()
		fillEmptySpaces()
		printBoard()
	}

	fmt.Println("Final Board:")
	printBoard()
}



func printBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func randomColor() int {
	return rand.Intn(5) + 1
}

func eliminate() bool {
	var eliminated bool

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 0 {
				continue
			}

			if eliminateHorizontal(i, j) || eliminateVertical(i, j) {
				eliminated = true
			}
		}
	}

	return eliminated
}

func eliminateHorizontal(x, y int) bool {
	count := 1
	for i := y + 1; i < boardSize; i++ {
		if board[x][i] == board[x][y] {
			count++
		} else {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if board[x][i] == board[x][y] {
			count++
		} else {
			break
		}
	}
	if count >= 3 {
		for i := y - count + 1; i <= y+count-1; i++ {
			board[x][i] = 0
		}
		return true
	}

	return false
}

func eliminateVertical(x, y int) bool {
	count := 1
	for i := x + 1; i < boardSize; i++ {
		if board[i][y] == board[x][y] {
			count++
		} else {
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == board[x][y] {
			count++
		} else {
			break
		}
	}
	if count >= 3 {
		for i := x - count + 1; i <= x+count-1; i++ {
			board[i][y] = 0
		}
		return true
	}

	return false
}

func fillEmptySpaces() {
	for i := boardSize - 1; i > 0; i-- {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 0 {
				for k := i - 1; k >= 0; k-- {
					if board[k][j] != 0 {
						//board
						board[i][j] = board[k][j]
						board[k][j] = 0
						break
					}
				}
			}
		}
	}

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 0 {
				board[i][j] = randomColor()
			}
		}
	}
}


func initializeBoard() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = randomColor()
		}
	}
	fmt.Println("Initial Board:")
	printBoard()
}


//////


func eliminate2(board [][]int) bool {
	boardSize := len(board)

	// 标记需要消除的格子
	var marked [][]bool
	for i := 0; i < boardSize; i++ {
		row := make([]bool, boardSize)
		marked = append(marked, row)
	}

	// 水平方向消除
	for i := 0; i < boardSize; i++ {
		j := 0
		for j < boardSize {
			if board[i][j] == 0 {
				j++
				continue
			}

			color := board[i][j]
			count := 1
			for k := j + 1; k < boardSize; k++ {
				if board[i][k] == color {
					count++
				} else {
					break
				}
			}

			if count >= 3 {
				for k := j; k < j+count; k++ {
					marked[i][k] = true
				}
			}

			j += count
		}
	}

	// 垂直方向消除
	for j := 0; j < boardSize; j++ {
		i := 0
		for i < boardSize {
			if board[i][j] == 0 {
				i++
				continue
			}

			color := board[i][j]
			count := 1
			for k := i + 1; k < boardSize; k++ {
				if board[k][j] == color {
					count++
				} else {
					break
				}
			}

			if count >= 3 {
				for k := i; k < i+count; k++ {
					marked[k][j] = true
				}
			}

			i += count
		}
	}

	// L 形消除
	for i := 0; i < boardSize-2; i++ {
		for j := 0; j < boardSize-1; j++ {
			if board[i][j] != 0 && board[i][j] == board[i+1][j] && board[i+2][j+1] == board[i+1][j] && board[i+2][j] == board[i+1][j] {
				marked[i][j] = true
				marked[i+1][j] = true
				marked[i+2][j] = true
				marked[i+2][j+1] = true
			}
		}
	}

	// T 形消除
	for i := 0; i < boardSize-2; i++ {
		for j := 1; j < boardSize-1; j++ {
			if board[i][j] != 0 && board[i][j] == board[i+1][j-1] && board[i+1][j] == board[i+2][j] && board[i+1][j+1] == board[i+1][j] {
				marked[i][j] = true
				marked[i+1][j] = true
				marked[i+1][j-1] = true
				marked[i+1][j+1] = true
				marked[i+2][j] = true
			}
		}
	}

	// 如果没有标记的格子，算法
	noMarked := true
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if marked[i][j] {
				noMarked = false
				board[i][j] = 0
			}
		}
	}

	if noMarked {
		return false
	}

	// 下落
	for j := 0; j < boardSize; j++ {
		nextEmptyRow := boardSize - 1
		for i := boardSize - 1; i >= 0; i-- {
			if board[i][j] != 0 {
				board[nextEmptyRow][j] = board[i][j]
				nextEmptyRow--
			}
		}
		for i := nextEmptyRow; i >= 0; i-- {
			board[i][j] = 0
		}
	}

	return true
}




