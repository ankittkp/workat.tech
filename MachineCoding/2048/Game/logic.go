package Game

import (
	"math/rand"
	"time"
)

func (b *board) AddNewElementAtBoard() {
	vacantPosition, position := 0, 0
	x, value := GetRandomElement()
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if b.grid[i][j] == 0 {
				vacantPosition++
			}
		}
	}
	elementCount := x%vacantPosition + 1
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if b.grid[i][j] == 0 {
				position++
				if position == elementCount {
					b.row = i
					b.col = j
					b.grid[i][j] = value
					return
				}
			}
		}
	}
	return
}

func GetRandomElement() (int, int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	value := r1.Int() % COMBINEDCHANCES
	if value <= TWOSCHANCE {
		value = 2
	} else {
		value = 4
	}
	return r1.Int(), value
}

func SlideMatrix(elems []int) []int {
	nonVacantPosition := make([]int, 0)
	for i := 0; i < COLUMNS; i++ {
		if elems[i] != 0 {
			nonVacantPosition = append(nonVacantPosition, elems[i])
		}
	}
	remaining := COLUMNS - len(nonVacantPosition)
	for i := 0; i < remaining; i++ {
		nonVacantPosition = append(nonVacantPosition, 0)
	}
	return mergeElements(nonVacantPosition)
}

func (b *board) ReverseMatrix() {
	for i := 0; i < ROWS; i++ {
		b.grid[i] = ReverseRow(b.grid[i])
	}
}

func (b *board) TransposeMatrix() {
	ans := make([][]int, 0)
	for i := 0; i < ROWS; i++ {
		ans = append(ans, make([]int, COLUMNS))
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			ans[i][j] = b.grid[COLUMNS-j-1][i]
		}
	}
	b.grid = ans
}

func (b *board) ReverseRows() {
	ans := make([][]int, 0)
	for i := 0; i < ROWS; i++ {
		ans = append(ans, make([]int, COLUMNS))
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			ans[ROWS-i-1][j] = b.grid[i][j]
		}
	}
	b.grid = ans
}

func ReverseRow(arr []int) []int {
	ans := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, arr[i])
	}
	return ans
}

func mergeElements(arr []int) []int {
	newArr := make([]int, len(arr))
	newArr[0] = arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == newArr[index] {
			newArr[index] += arr[i]
		} else {
			index++
			newArr[index] = arr[i]
		}
	}
	return newArr
}