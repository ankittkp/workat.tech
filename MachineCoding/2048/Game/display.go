package Game

import (
	"fmt"
	"github.com/fatih/color"
)

func (b *board) GameShow() {
	d := color.New(color.FgBlue, color.Bold)
	fmt.Println(CLRSCR)
	for i := 0; i < len(b.grid); i++ {
		printHorizontal()
		fmt.Printf("|")
		for j := 0; j < len(b.grid[0]); j++ {
			fmt.Printf("%3s", "")
			if b.grid[i][j] == 0 {
				fmt.Printf("%-6s|", "")
			} else {
				if i == b.row && j == b.col {
					d.Printf("%-6d|", b.grid[i][j])
				} else {
					fmt.Printf("%-6d|", b.grid[i][j])
				}
			}
		}
		fmt.Printf("%4s", "")
		fmt.Println()
	}
	printHorizontal()
}
func printHorizontal() {
	for i := 0; i < 40; i++ {
		fmt.Print("-")
	}
	fmt.Printf("\n")
}
