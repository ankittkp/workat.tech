package Game

import (
	"errors"
	"github.com/eiannone/keyboard"
	"log"
)
var EndGame = errors.New("GameOver")
func GetKeyboardInput() (Direction, error) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	char, key, err := keyboard.GetKey()
	ans := int(char)
	if ans == 0 {
		ans = int(key)
	}
	if DebugLogLevel {
		log.Printf("the key is: %v \n", ans)
	}
	if err != nil {
		return CANCEL, err
	}
	switch ans {
	case 119, 65517, 107:
		return UP, nil
	case 97, 65515, 104:
		return LEFT, nil
	case 115, 65516, 106:
		return DOWN, nil
	case 100, 65514, 108:
		return RIGHT, nil
	case 3:
		return CANCEL, EndGame
	}
	return CANCEL, nil
}

func (b *board) Input() {
	var dir Direction
	dir, err := GetKeyboardInput()
	if err != nil {
		if errors.Is(err, EndGame) {
			b.done = true
			return
		} else {
			log.Fatalf("Error while taking input: %v", err)
			return
		}
	}
	if DebugLogLevel {
		log.Printf("the dir is: %v \n", dir)
	}

	if dir == CANCEL {
		b.Input()
	}
	b.move(dir)
}
