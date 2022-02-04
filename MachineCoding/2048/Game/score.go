package Game

import (
	"io/ioutil"
	"strconv"
)

func (b *board) GameScore() (int, int) {
	total := 0
	currentScore := 0
	highScore := 0
	data, err := ioutil.ReadFile("highscore.txt")
	existingHighScore, _ := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}
	grid := b.grid
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			total += grid[i][j]
			currentScore = max(currentScore, grid[i][j])
		}
	}
	if currentScore > existingHighScore {
		highScore = currentScore
		err = ioutil.WriteFile("highscore.txt", []byte(strconv.Itoa(currentScore)), 0644)
		if err != nil {
			panic(err)
		}
	} else{
		highScore = existingHighScore
	}
	return currentScore, highScore
}

func max(one int, two int) int {
	if one > two {
		return one
	}
	return two
}

func (b *board) GameOver() bool {
	vacantPosition := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if b.grid[i][j] == 0 {
				vacantPosition++
			}
		}
	}
	return vacantPosition == 0 || b.done
}
