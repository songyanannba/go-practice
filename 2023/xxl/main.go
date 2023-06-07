//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//const boardSize = 7
//const NUM = 4
//
//
//
//var board [boardSize][boardSize]int
//
///*该算法实现了一个简单的消消乐游戏，包括棋盘的初始化、打印、连续消除和填充空位等操作。
//其中 `eliminate` 函数实现了横向、竖向和斜向的连续消除规则，该函数返回一个布尔值，表示是否进行了消除。
//`fillEmptySpaces` 函数实现了填充空位的操作，即将上方非空格子的颜色向下移动，同时在第一行随机生成新的颜色。
//
//该算法中使用了一个二维数组来表示棋盘，并且使用了 Golang 内置的随机数生成器来随机生成颜色。
//如果需要提高算法的复杂度，可以考虑使用其他更加复杂的算法来生成随机颜色，并且对连续消除的规则进行优化。*/
//
//func main() {
//	var size = 0
//	rand.Seed(time.Now().UnixNano())
//	initBoard()
//	printBoard()
//	fmt.Println("Start Eliminating...")
//
//
//
//	for {
//		eliminated := eliminateAll()
//		if !eliminated {
//			fmt.Println("*** eliminated false")
//			break
//		} else {
//			fmt.Println("*** eliminated true")
//		}
//
//		fillEmptySpaces()
//		fmt.Println("size == " , size)
//		printBoard11()
//		size++
//		time.Sleep(500 * time.Millisecond)
//	}
//
//	fmt.Println("Game Over")
//}
//
//func initBoard() {
//	for i := 0; i < boardSize; i++ {
//		for j := 0; j < boardSize; j++ {
//			board[i][j] = randomColor()
//		}
//	}
//}
//
//func printBoard11() {
//	for i := 0; i < boardSize; i++ {
//		for j := 0; j < boardSize; j++ {
//			fmt.Print(board[i][j], " ")
//		}
//		fmt.Println()
//	}
//	fmt.Println("== 完成 printBoard111")
//}
//
//func printBoard22() {
//	for i := 0; i < boardSize; i++ {
//		for j := 0; j < boardSize; j++ {
//			fmt.Print(board[i][j], " ")
//		}
//		fmt.Println()
//	}
//	fmt.Println("== 补充数据 printBoard22")
//}
//
//func randomColor() int {
//	return rand.Intn(5) + 1
//}
//
//func printBoard() {
//	for i := 0; i < boardSize; i++ {
//		for j := 0; j < boardSize; j++ {
//			fmt.Print(board[i][j], " ")
//		}
//		fmt.Println()
//	}
//	fmt.Println("== 完成 printBoard")
//}
//
//func eliminateAll() bool {
//	eliminated := false
//	for i := 0; i < boardSize; i++ {
//		for j := 0; j < boardSize; j++ {
//			if eliminate(i, j) {
//				eliminated = true
//			}
//		}
//	}
//
//
//	fmt.Println("== 完成 eliminateAll == ")
//	printBoard()
//
//	return eliminated
//}
//
//func eliminate(x, y int) bool {
//	if board[x][y] == 0 {
//		return false
//	}
//
//	color := board[x][y]
//
//	// Check horizontal
//	count := 0
//	for i := y; i < boardSize; i++ {
//		if board[x][i] == color {
//			count++
//		} else {
//			break
//		}
//	}
//	if count >= NUM {
//		for i := y; i < y+count; i++ {
//			board[x][i] = 0
//		}
//		fmt.Println("horizontal true", color)
//		return true
//	}
//
//	// Check vertical
//	count = 0
//	for i := x; i < boardSize; i++ {
//		if board[i][y] == color {
//			count++
//		} else {
//			break
//		}
//	}
//	if count >= NUM {
//		for i := x; i < x+count; i++ {
//			board[i][y] = 0
//		}
//		fmt.Println("vertical true" ,color)
//		return true
//	}
//
//	// Check diagonal
//	/*count = 0
//	for i, j := x, y; i < boardSize && j < boardSize; i, j = i+1, j+1 {
//		if board[i][j] == color {
//			count++
//		} else {
//			break
//		}
//	}
//	if count >= 3 {
//		for i, j := x, y; i < x+count && j < y+count; i, j = i+1, j+1 {
//			board[i][j] = 0
//		}
//		fmt.Println("diagonal  true", color)
//		return true
//	}
//
//	count = 0
//	for i, j := x, y; i >= 0 && j < boardSize; i, j = i-1, j+1 {
//		if board[i][j] == color {
//			count++
//		} else {
//			//fmt.Println("diagonal   board[i][j] true 11")
//			break
//		}
//	}
//	if count >= 3 {
//		for i, j := x, y; i > x-count && j < y+count; i, j = i-1, j+1 {
//			board[i][j] = 0
//		}
//		fmt.Println("diagonal   true 22" ,color)
//		return true
//	}*/
//
//	return false
//}
//
//func fillEmptySpaces() {
//	for i := boardSize - 1; i > 0; i-- {
//		for j := 0; j < boardSize; j++ {
//			if board[i][j] == 0 {
//				for k := i - 1; k >= 0; k-- {
//					if board[k][j] != 0 {
//						board[i][j] = board[k][j]
//						board[k][j] = 0
//						break
//					}
//				}
//			}
//		}
//	}
//
//	printBoard22()
//
//	for i := 0; i < boardSize; i++ {
//		if board[0][i] == 0 {
//			board[0][i] = randomColor()
//		}
//	}
//}
