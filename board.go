package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

type Board struct {
	dimension          int
	Cells              [][]int
	PositionOfFreeCell []int
	PossibleMoves      []rune
}

func NewBoard(dimension int) *Board {
	b := &Board{dimension: dimension}
	b.Cells = make([][]int, dimension)
	counter := 1
	for i := range b.Cells {
		b.Cells[i] = make([]int, dimension)
		for j := range b.Cells[i] {
			b.Cells[i][j] = counter % (dimension * dimension)
			counter++
		}
	}
	return b
}

func (b *Board) print() {
	var str string
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			fmt.Printf("%2d ", b.Cells[i][j])
		}
		fmt.Println()
	}
	fmt.Println(str)
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
	b.PossibleMoves = []rune{}

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

func (b *Board) MakeMove(move rune) {
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
	randomIdx := rand.Intn(len(b.PossibleMoves))
	b.MakeMove(b.PossibleMoves[randomIdx])
	b.UpdatePositionOfFreeCell()
}

func (b *Board) Shuffle() {
	for _ = range b.dimension * ShuffleCount {
		b.MakeRandomMove()
	}
}

func (b *Board) IsEndGame() bool {
	tempBoard := NewBoard(b.dimension)
	return b.Equal(tempBoard)
}

func (b *Board) ClearAndShow() {
	//cmd := exec.Command("cmd", "/c", "cls")
	//
	//cmd.Stdout = os.Stdout
	//err := cmd.Run()
	//if err != nil {
	//	return
	//}

	b.print()
}

func (b *Board) StartGame() {
	b.Shuffle()
	b.ClearAndShow()
	counter := 0

	start := time.Now()
	for !b.IsEndGame() {
		b.UpdatePossibleMoves()
		b.UpdatePositionOfFreeCell()

		var char rune
		_, _ = fmt.Scanf("%c", &char)

		if slices.Contains(b.PossibleMoves, char) {
			b.MakeMove(char)
			counter++
		} else {
			strMoves := make([]string, len(b.PossibleMoves))
			for i, move := range b.PossibleMoves {
				strMoves[i] = fmt.Sprintf("%c", move)
			}
			fmt.Printf("You can only use %v moves\n", strMoves)
		}

		b.ClearAndShow()
	}

	fmt.Printf("Congratulations! You won the game in %d moves!\n", counter)
	fmt.Printf("Time:  %s", time.Since(start))
}
