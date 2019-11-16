package tictactoe

import (
	"fmt"
	"strings"
)

// Row defines a row on the board.
type Row uint8

// The following rows exist:
const (
	RowA Row = iota
	RowB
	RowC
)

// Col defines a column on the board.
type Col uint8

// The followinf columns exist:
const (
	Col1 Col = iota
	Col2
	Col3
)

// Square represents a square on the board.
type Square uint8

// The following square exist:
const (
	A1 Square = iota
	A2
	A3
	B1
	B2
	B3
	C1
	C2
	C3
)

// Row returns the row of this squre.
func (sq Square) Row() Row {
	switch sq {
	case A1, A2, A3:
		return RowA
	case B1, B2, B3:
		return RowB
	case C1, C2, C3:
		return RowC
	default:
		panic(fmt.Errorf("invalid square: %v", sq))
	}
}

// Col returns the column of this squre.
func (sq Square) Col() Col {
	switch sq {
	case A1, B1, C1:
		return Col1
	case A2, B2, C2:
		return Col2
	case A3, B3, C3:
		return Col3
	default:
		panic(fmt.Errorf("invalid square: %v", sq))
	}
}

// Player is defines as:
type Player rune

// Two players exist. The NoPlayer is used to denote a vacant square.
const (
	Player1  = 'X'
	Player2  = 'O'
	NoPlayer = ' '
)

var (
	errDraw   = fmt.Errorf("GAME OVER")
	errWinner = fmt.Errorf("YOU WIN")
	errLooser = fmt.Errorf("YOU LOOSE")
)

// IsGameOver returns true if the given error indicates that the game is over.
func IsGameOver(err error) bool {
	return err == errDraw || err == errWinner || err == errLooser
}

// Board represents a tic-tac-toe game board.
type Board struct {
	Player1 []Square
	Player2 []Square
	result  error
}

// NewBoard returns an initialized Board.
func NewBoard() *Board {
	return &Board{
		Player1: []Square{},
		Player2: []Square{},
	}
}

// String returns a ASCII-art board representation.
func (b Board) String() string {
	return fmt.Sprintf(
		"\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n",
		string(b.OwnerOf(A1)), string(b.OwnerOf(A2)), string(b.OwnerOf(A3)),
		string(b.OwnerOf(B1)), string(b.OwnerOf(B2)), string(b.OwnerOf(B3)),
		string(b.OwnerOf(C1)), string(b.OwnerOf(C2)), string(b.OwnerOf(C3)),
	)
}

// NextPlayer returns the player who has the next move.
func (b Board) NextPlayer() Player {
	if len(b.Player1) == len(b.Player2) {
		return Player1
	}
	return Player2
}

// OwnerOf returns the Player who has a piece on the given square or NoPlayer.
func (b Board) OwnerOf(sq Square) Player {
	for i := range b.Player1 {
		if b.Player1[i] == sq {
			return Player1
		}
	}
	for i := range b.Player2 {
		if b.Player2[i] == sq {
			return Player2
		}
	}
	return NoPlayer
}

// ParseMove parses the given string and plays the parsed move.
func (b *Board) ParseMove(move string) error {
	move = strings.ToUpper(move)
	switch move {
	case "A1":
		return b.Set(A1)
	case "A2":
		return b.Set(A2)
	case "A3":
		return b.Set(A3)
	case "B1":
		return b.Set(B1)
	case "B2":
		return b.Set(B2)
	case "B3":
		return b.Set(B3)
	case "C1":
		return b.Set(C1)
	case "C2":
		return b.Set(C2)
	case "C3":
		return b.Set(C3)
	}
	return fmt.Errorf("can not parse: %s", move)
}

// Set makes a move by setting the next sqare.
func (b *Board) Set(sq Square) error {
	nextPlayer := b.NextPlayer()
	return b.setSquare(nextPlayer, sq)
}

func (b *Board) setSquare(p Player, sq Square) error {
	if b.result != nil {
		return b.result
	}
	switch sq {
	case A1, A2, A3, B1, B2, B3, C1, C2, C3:
	default:
		return fmt.Errorf("invalid square: %v", sq)
	}
	owner := b.OwnerOf(sq)
	if owner != NoPlayer {
		return fmt.Errorf("sqaure is occupied by %s", string(owner))
	}
	switch p {
	case Player1:
		b.Player1 = append(b.Player1, sq)
	case Player2:
		b.Player2 = append(b.Player2, sq)
	default:
		return fmt.Errorf("invalid player: %v", p)
	}

	if b.isWinner(p) {
		switch p {
		// default is caught above
		case Player1:
			b.result = errWinner
		case Player2:
			b.result = errLooser
		}
		return b.result
	}

	if len(b.Player1)+len(b.Player2) == 9 {
		b.result = errDraw
		return b.result
	}

	return nil
}

func (b Board) isWinner(p Player) bool {
	fields := []Square{}
	switch p {
	case Player1:
		fields = b.Player1
	case Player2:
		fields = b.Player2
	default:
		panic(fmt.Errorf("invalid player: %v", p))
	}

	if len(fields) < 3 {
		return false
	}

	rows := map[Row]int{
		RowA: 0,
		RowB: 0,
		RowC: 0,
	}
	cols := map[Col]int{
		Col1: 0,
		Col2: 0,
		Col3: 0,
	}
	diag1 := 0
	diag2 := 0

	for i := range fields {
		rows[fields[i].Row()]++
		cols[fields[i].Col()]++

		if fields[i] == A1 || fields[i] == B2 || fields[i] == C3 {
			diag1++
		}
		if fields[i] == A3 || fields[i] == B2 || fields[i] == C1 {
			diag2++
		}
	}

	for i := range rows {
		if rows[i] == 3 {
			return true
		}
	}
	for i := range cols {
		if cols[i] == 3 {
			return true
		}
	}

	if diag1 == 3 || diag2 == 3 {
		return true
	}
	return false
}
