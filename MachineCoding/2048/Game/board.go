package Game

var DebugLogLevel bool
const (
	ROWS = 4
	COLUMNS = 4
	CLRSCR = "\033[H\033[2J"
	COMBINEDCHANCES = 100
	TWOSCHANCE = 80
)
type board struct {
	grid [][]int
	done   bool
	row int
	col int
}
type Board interface {
	AddNewElementAtBoard()
	Input()
	GameScore() (int, int)
	GameShow()
	GameOver() bool
}

func CreateNewBoard() Board {
	grid := make([][]int, 0)
	for i := 0; i < ROWS; i++ {
		grid = append(grid, make([]int, COLUMNS))
	}
	return &board{
		grid: grid,
	}
}
