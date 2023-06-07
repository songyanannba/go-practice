package main

import (
	"fmt"
	"math/rand"
)

const (
	rows = 7
	cols = 7
)

// 检查一个方块是否在格子内
func inBounds(x, y int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

// 检查一个方块是否与其相邻的方块相同
func isAdjacentMatch(grid [][]int, x, y int) bool {
	if !inBounds(x, y) {
		return false
	}

	val := grid[x][y]

	// 检查上下左右四个方向
	if inBounds(x-1, y) && grid[x-1][y] == val {
		return true
	}
	if inBounds(x+1, y) && grid[x+1][y] == val {
		return true
	}
	if inBounds(x, y-1) && grid[x][y-1] == val {
		return true
	}
	if inBounds(x, y+1) && grid[x][y+1] == val {
		return true
	}

	// 如果包含T或L的情况，需要检查对角线上的两个方向
	if (x == 0 && y == 0) || (x == rows-1 && y == 0) || (x == 0 && y == cols-1) || (x == rows-1 && y == cols-1) {
		return false
	}
	if inBounds(x-1, y-1) && grid[x-1][y-1] == val {
		if inBounds(x-1, y+1) && grid[x-1][y+1] == val {
			return true
		}
		if inBounds(x+1, y-1) && grid[x+1][y-1] == val {
			return true
		}
	}
	if inBounds(x+1, y+1) && grid[x+1][y+1] == val {
		if inBounds(x-1, y+1) && grid[x-1][y+1] == val {
			return true
		}
		if inBounds(x+1, y-1) && grid[x+1][y-1] == val {
			return true
		}
	}

	return false
}

// 检查是否有任何相邻的方块可以被消除
func hasMatches(grid [][]int) bool {
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if isAdjacentMatch(grid, x, y) {
				return true
			}
		}
	}
	return false
}

// 从格子中消除匹配的方块
func removeMatches(grid [][]int) int {
	matches := 0

	// 标记要消除的方块
	marked := make([][]bool, rows)
	for i := range marked {
		marked[i] = make([]bool, cols)
	}
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if isAdjacentMatch(grid, x, y) {
				marked[x][y] = true
			}
		}
	}

	// 消除标记的方块，并将 方块移动到空出的位置
	for y := 0; y < cols; y++ {
		dst := rows - 1
		for src := rows - 1; src >= 0; src-- {
			if !marked[src][y] {
				grid[dst][y] = grid[src][y]
				dst--
			}
		}
		// 将顶部的空出的位置填充为随机方块
		for dst >= 0 {
			grid[dst][y] = randBlock()
			dst--
		}
	}

	// 计算消除的方块数量
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if marked[x][y] {
				matches++
			}
		}
	}

	return matches
}

// 随机生成一个方块
func randBlock() int {
	return rand.Intn(5) + 1
}

// 该算法的基本思路是从左到右、从上到下依次遍历每个方块，检查其是否与相邻的方块相同，如果是则标记为要消除的方块。
//然后，从底部开始，将非标记的方块依次移动到空出的位置，同时将顶部的空出的位置填充为随机方块。
//最后，计算消除的方块数量，直到没有匹配的方块。如果需要考虑包含T和L的情况，则需要在检查相邻的方块时，同时检查对角线上的方块。

// 主函数
func main() {
	// 初始化一个7x7的格子，并随机填充方块
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = randBlock()
		}
	}

	// 持续消除匹配的方块，直到没有匹配的方块
	for hasMatches(grid) {
		removeMatches(grid)
	}

	fmt.Println("完成！")
}