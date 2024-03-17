package main

import "math/rand"

type Board struct {
	dimension          int
	Cells              [][]int
	PositionOfFreeCell []int
	PossibleMoves      []string
}

func NewBoard(dimension int) *Board {
	b := &Board{dimension: dimension}
	b.Cells = make([][]int, dimension)
	for i := range b.Cells {
		b.Cells[i] = make([]int, dimension)
	}
	return b
}

func (b *Board) print() string {
	var str string
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			str += " " + string(b.Cells[i][j])
		}
		str += "\n"
	}
	return str
}

func (b *Board) Equal(other *Board) bool {
	if b.dimension != other.dimension {
		return false
	}

	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.Cells[i][j] != other.Cells[i][j] {
				return false
			}
		}
	}

	return true
}

func (b *Board) IsFreeCell(i, j int) bool {
	return b.Cells[i][j] == 0
}

func (b *Board) UpdatePositionOfFreeCell() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.IsFreeCell(i, j) {
				b.PositionOfFreeCell = []int{i, j}
			}
		}
	}
}

func (b *Board) UpdatePossibleMoves() {
	b.PossibleMoves = []string{}

	b.UpdatePositionOfFreeCell()
	i, j := b.PositionOfFreeCell[0], b.PositionOfFreeCell[1]

	if i > 0 {
		b.PossibleMoves = append(b.PossibleMoves, DOWN)
	}

	if i != b.dimension-1 {
		b.PossibleMoves = append(b.PossibleMoves, UP)
	}

	if j > 0 {
		b.PossibleMoves = append(b.PossibleMoves, RIGHT)
	}

	if j != b.dimension-1 {
		b.PossibleMoves = append(b.PossibleMoves, LEFT)
	}
}

func (b *Board) MoveUp() {
	i, j := b.PositionOfFreeCell[0], b.PositionOfFreeCell[1]
	b.Cells[i][j], b.Cells[i+1][j] = b.Cells[i+1][j], b.Cells[i][j]
	b.UpdatePositionOfFreeCell()
}

func (b *Board) MoveDown() {
	i, j := b.PositionOfFreeCell[0], b.PositionOfFreeCell[1]
	b.Cells[i][j], b.Cells[i-1][j] = b.Cells[i-1][j], b.Cells[i][j]
	b.UpdatePositionOfFreeCell()
}

func (b *Board) MoveLeft() {
	i, j := b.PositionOfFreeCell[0], b.PositionOfFreeCell[1]
	b.Cells[i][j], b.Cells[i][j+1] = b.Cells[i][j+1], b.Cells[i][j]
	b.UpdatePositionOfFreeCell()
}

func (b *Board) MoveRight() {
	i, j := b.PositionOfFreeCell[0], b.PositionOfFreeCell[1]
	b.Cells[i][j], b.Cells[i][j-1] = b.Cells[i][j-1], b.Cells[i][j]
	b.UpdatePositionOfFreeCell()
}

func (b *Board) MakeMove(move string) {
	switch move {
	case UP:
		b.MoveUp()
	case DOWN:
		b.MoveDown()
	case LEFT:
		b.MoveLeft()
	case RIGHT:
		b.MoveRight()
	}
}

func (b *Board) MakeRandomMove() {
	b.UpdatePossibleMoves()
	randomIdx := rand.Int31(0, len(b.PossibleMoves))
	b.MakeMove(b.PossibleMoves[randomIdx])
	b.UpdatePositionOfFreeCell()
}

func (b *Board) Shuffle() {
	for i := 0; i < ShuffleCount; i++ {
		b.MakeRandomMove()
	}
}

func (b *Board) IsEndGame() bool {
	tempBoard := NewBoard(b.dimension)
	return b.Equal(tempBoard)
}

func (b *Board) StartGame() {
	b.Shuffle()
}
