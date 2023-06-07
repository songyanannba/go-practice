package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ROWS      = 7
	COLS      = 7
	MIN_MATCH = 3
)

type Board struct {
	cells [][]int
}

func NewBoard() *Board {
	cells := make([][]int, ROWS)
	for i := 0; i < ROWS; i++ {
		cells[i] = make([]int, COLS)
	}
	return &Board{cells}
}

func (b *Board) Init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			b.cells[i][j] = rand.Intn(5) + 1
		}
	}
}

func (b *Board) Print() {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			fmt.Printf("%d ", b.cells[i][j])
		}
		fmt.Println("---")
	}
	fmt.Println("===")
}

func (b *Board) IsMatch(row, col int) bool {
	if b.cells[row][col] == 0 {
		return false
	}

	matchCount := 1
	color := b.cells[row][col]

	// 上方向
	for i := row - 1; i >= 0; i-- {
		if b.cells[i][col] == color {
			matchCount++
		} else {
			break
		}
	}

	// 下方向
	for i := row + 1; i < ROWS; i++ {
		if b.cells[i][col] == color {
			matchCount++
		} else {
			break
		}
	}

	// 左方向
	for j := col - 1; j >= 0; j-- {
		if b.cells[row][j] == color {
			matchCount++
		} else {
			break
		}
	}

	// 右方向
	for j := col + 1; j < COLS; j++ {
		if b.cells[row][j] == color {
			matchCount++
		} else {
			break
		}
	}

	// T 形状判断
	if matchCount < MIN_MATCH {
		return false
	}

	if matchCount >= MIN_MATCH {
		if row > 0 && row < ROWS-1 && col > 0 && col < COLS-1 {
			if b.cells[row-1][col] == color && b.cells[row+1][col] == color && b.cells[row][col-1] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
			if b.cells[row-1][col] == color && b.cells[row][col-1] == color && b.cells[row][col+1] == color {

				fmt.Println("color ++= " , color)
				return true
			}
			//if b.cells[row+1][col] == color &&
			if b.cells[row][col-1] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
			if b.cells[row-1][col] == color && b.cells[row+1][col] == color && b.cells[row][col-1] == color {

				fmt.Println("color ++= " , color)
				return true
			}
			if b.cells[row-1][col] == color && b.cells[row+1][col] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		return false
	}

	// L 形状判断
	if matchCount > MIN_MATCH+1 {
		if row > 0 && col > 0 && col < COLS-1 {
			if b.cells[row-1][col] == color && b.cells[row][col-1] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		if row < ROWS-1 && col > 0 && col < COLS-1 {
			if b.cells[row+1][col] == color && b.cells[row][col-1] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		if col > 1 {
			if b.cells[row][col-2] == color && b.cells[row][col-1] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		if col < COLS-2 {
			if b.cells[row][col+2] == color && b.cells[row][col+1] == color && b.cells[row][col-1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		if row > 1 && col > 0 {
			if b.cells[row-2][col] == color && b.cells[row-1][col] == color && b.cells[row][col-1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		if row > 1 && col < COLS-1 {
			if b.cells[row-2][col] == color && b.cells[row-1][col] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		if row < ROWS-2 && col > 0 {
			if b.cells[row+2][col] == color && b.cells[row+1][col] == color && b.cells[row][col-1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		if row < ROWS-2 && col < COLS-1 {
			if b.cells[row+2][col] == color && b.cells[row+1][col] == color && b.cells[row][col+1] == color {
				fmt.Println("color ++= " , color)
				return true
			}
		}
		return false
	}

	return false
}

func (b *Board) RemoveMatch() bool {
	removed := false
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if b.IsMatch(i, j) {

				b.cells[i][j] = 0
				removed = true
			}
		}
	}
	return removed
}

func (b *Board) Collapse() {
	for j := 0; j < COLS; j++ {
		for i := ROWS - 1; i >= 0; i-- {
			if b.cells[i][j] == 0 {
				for k := i - 1; k >= 0; k-- {
					if b.cells[k][j] != 0 {
						b.cells[i][j] = b.cells[k][j]
						b.cells[k][j] = 0
						break
					}
				}
			}
		}
	}
}

func (b *Board) Play() {
	b.Init()
	for {
		b.Print()
		if !b.RemoveMatch() {
			break
		}
		b.Collapse()
		time.Sleep(time.Second * 10)
	}
}
// 在这个版本的代码中，我们加入了 L 形和 T 形匹配的检查，并增加了相应的测试代码。
//另外，我们也加入了支持连续玩法的代码，游戏结束的条件是不再有能够消除的连续块。运行上述代码，你应该能看到如下类似的输出：



type Tag struct {
	Id       int
	Name     string
	Include  []string `json:"-"`
	Multiple int

	X int
	Y int

	IsLine     bool `json:"-"`
	IsPayTable bool `json:"-"`
	IsWild     bool `json:"-"`
	IsSingle   bool `json:"-"` // 是否单出
	IsJackpot  bool `json:"-"`
	ISLock     bool `json:"-"` // 是否锁定
}

func main() {

	in := []int{1, 2, 3, 4, 5}
	out := make([]*int, 0)
	for _, v := range in {
		//v := v  打开注释即正确
		out = append(out, &v)
	}

	fmt.Println("res:", *out[0], *out[1], *out[2])



/*	t := Tag{
		Id:         1,
		Name:       "fff",
		Include:    nil,
		Multiple:  int(0.2),
		X:          0,
		Y:          0,
		tzIsLine:     false,
		IsPayTable: false,
		IsWild:     true,
		IsSingle:   false,
		IsJackpot:  false,
		ISLock:     false,
	}

	marshal, _ := json.Marshal(t)

	fmt.Println(string(marshal))
	fmt.Println(string(11))*/
	//json.Unmarshal()

	//board := NewBoard()
	//board.Play()
}