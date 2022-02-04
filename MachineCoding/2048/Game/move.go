package Game

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	CANCEL
)

func (b *board) move(dir Direction) {
	switch dir {
	case LEFT:
		b.moveLeft()
	case RIGHT:
		b.moveRight()
	case DOWN:
		b.moveDown()
	case UP:
		b.moveUp()
	}
}

func (b *board) moveLeft() {
	for i := 0; i < ROWS; i++ {
		old := b.grid[i]
		b.grid[i] = SlideMatrix(old)
	}
}

func (b *board) moveUp() {
	b.ReverseRows()
	b.moveDown()
	b.ReverseRows()
}

func (b *board) moveDown() {
	b.TransposeMatrix()
	b.moveLeft()
	b.TransposeMatrix()
	b.TransposeMatrix()
	b.TransposeMatrix()
}

func (b *board) moveRight() {
	b.ReverseMatrix()
	b.moveLeft()
	b.ReverseMatrix()
}