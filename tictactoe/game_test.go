package tictactoe

import (
	"fmt"
	"testing"
)

var (
	// boardTemplate serves as template string for formatting board positions in testing
	boardTemplate = "\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n"

	// emptyBoard renders above template as an empty board with zero moves made
	emptyBoard = fmt.Sprintf(
		boardTemplate,
		" ", " ", " ",
		" ", " ", " ",
		" ", " ", " ",
	)
)

func TestNewBoard(t *testing.T) {
	// initialize a board
	b := NewBoard()

	// raise if the board doesn't render as an empty board
	if b.String() != emptyBoard {
		t.Errorf("expected %v, got %v", emptyBoard, b.String())
	} else {
		t.Log("Got expected: empty board")
		t.Log(b)
	}
}

func TestBoardPositions(t *testing.T) {
	// initialize a board
	b := NewBoard()

	// make the first move on A1
	b.Set(A1)

	// raise if the board doesn't render as a board where Player1 played A1
	boardSetup := fmt.Sprintf(
		boardTemplate,
		string(Player1), " ", " ",
		" ", " ", " ",
		" ", " ", " ",
	)
	if b.String() != boardSetup {
		t.Errorf("expected %v, got %v", boardSetup, b.String())
	} else {
		t.Log("Got expected: A1")
		t.Log(b)
	}

	// make the second move on C3
	b.Set(C3)

	// raise if the board doesn't render as a board where Player1 played A1 and Player2 playd C3
	boardSetup = fmt.Sprintf(
		boardTemplate,
		string(Player1), " ", " ",
		" ", " ", " ",
		" ", " ", string(Player2),
	)
	if b.String() != boardSetup {
		t.Errorf("expected %v, got %v", boardSetup, b.String())
	} else {
		t.Log("Got expected: A1, C3")
		t.Log(b)
	}
}

func TestIllegalMoves(t *testing.T) {
	// initialize a board
	b := NewBoard()

	// make first move on A1
	err := b.Set(A1)

	// raise if there was an error
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	// make the second move also on A1
	err = b.Set(A1)

	// raise if there was no error
	if err == nil {
		t.Errorf("expected != nil, got %v", err)
	}

	// make a move outside the board scope
	err = b.Set(Square(9))

	// raise if there was no error
	if err == nil {
		t.Errorf("expected != nil, got %v", err)
	}

	// make parsed move outside of board scope
	err = b.ParseMove("a0")

	// raise if there was no error
	if err == nil {
		t.Errorf("expected != nil, got %v", err)
	}
}

func TestWins(t *testing.T) {
	for _, wSet := range winningSets {
		var err error = nil
		b := NewBoard()
		for _, sq := range wSet {
			err = b.setSquare(Player1, sq)
		}
		if err != nil {
			if IsGameOver(err) {
				t.Log(b)
				t.Log("Got expected:", err)
			} else {
				t.Error(err)
			}
		}
	}
}

func TestLosses(t *testing.T) {
	for _, lSet := range winningSets {
		var err error = nil
		b := NewBoard()
		for _, sq := range lSet {
			err = b.setSquare(Player2, sq)
		}
		if err != nil {
			if IsGameOver(err) {
				t.Log(b)
				t.Log("Got expected:", err)
			} else {
				t.Error(err)
			}
		}
	}
}
