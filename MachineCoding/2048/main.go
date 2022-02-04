package main

import (
	"flag"
	"fmt"
	"github.com/jinxankit/workat.tech/MachineCoding/2048/Game"

)

var welcome = [...]string{
	"",
	"",
	"HOW TO PLAY: Use Arrow Keys or W, A, S, D or H, J, K, L to move the",
	"tiles. Tiles with the same number merge into one when they touch.",
	"Add them up to reach 2048 or more ??",
	"Press Any Key to Start and CTRL+C to Quit",
	"",
	"",
}

var logo = [...]string{
	"  .-----.   .----.      .---.    .-----.      ",
	" / ,-.   \\ /  ..  \\    / .  |   /  .-.  \\  ",
	" '-'  |  |.  /  \\  .  / /|  |  |   \\_.' /   ",
	"    .'  / |  |  '  | / / |  |_  /  .-. '.     ",
	"  .'  /__ '  \\  /  '/  '-'    ||  |   |  |   ",
	" |       | \\  `'  / `----|  |-' \\  '-'  /   ",
	"  -------'  `---''       `--'    `----''      ",
	"                                              ",
}

func main() {
	debug := flag.Bool("debug", false, "debugging flag")
	flag.Parse()
	if *debug {
		Game.DebugLogLevel = *debug
	}
	fmt.Printf("")
	for _, line := range welcome {
		fmt.Println(line)
	}
	for _, line := range logo {
		fmt.Println(line)
	}
	_, err := Game.GetKeyboardInput()
	if err != nil {
		fmt.Printf("Error while taking input to start the game: %v", err)
	}
	g := Game.CreateNewBoard()
	g.AddNewElementAtBoard()
	g.AddNewElementAtBoard()
	for true {
		if g.GameOver() {
			break
		}
		g.AddNewElementAtBoard()
		g.GameShow()
		g.Input()
	}
	fmt.Printf("\n\n**** Game Over **** \n")
	currentScore, highScore := g.GameScore()
	fmt.Printf("Current Score:    %v \n", currentScore)
	fmt.Printf("High Score: %v \n", highScore)
}
