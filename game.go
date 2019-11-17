package tictactoe

import (
	"context"
	"fmt"
	"log"
	"sync"

	game "github.com/baccenfutter/tictactoe/gen/game"
	backend "github.com/baccenfutter/tictactoe/tictactoe"
	uuid "github.com/satori/go.uuid"
)

var boards = sync.Map{}

// game service example implementation.
// The example methods log the requests and return zero values.
type gamesrvc struct {
	logger *log.Logger
}

// NewGame returns the game service implementation.
func NewGame(logger *log.Logger) game.Service {
	return &gamesrvc{logger}
}

// New initialize a new board
func (s *gamesrvc) New(ctx context.Context) (res *game.NewResult, err error) {
	res = &game.NewResult{}
	id := uuid.NewV4().String()
	boards.Store(id, backend.NewBoard())
	res.ID = id
	return
}

// Get obtains a board by its ID
func (s *gamesrvc) Get(ctx context.Context, p *game.GetPayload) (res *game.GetResult, err error) {
	res = &game.GetResult{}
	return
}

// Move implements move.
func (s *gamesrvc) Move(ctx context.Context, p *game.MovePayload) (res *game.MoveResult, err error) {
	res = &game.MoveResult{}
	// get board
	result, ok := boards.Load(p.Board)
	if !ok {
		err = game.MakeNotFound(fmt.Errorf("board does not exist: %s", p.Board))
		return
	}
	board := result.(*backend.Board)

	// make move
	err = board.ParseMove(p.Square)
	if err != nil {
		if backend.IsGameOver(err) {
			winner := err.Error()
			res.Winner = &winner
			err = nil
			boards.Delete(p.Board)
		} else {
			err = game.MakeBadRequest(err)
			return
		}
	}

	// make computer move
	backend.MakeOptimalMove(board)
	s.logger.Println(board)

	// prepare output
	a1 := string(board.OwnerOf(backend.A1))
	a2 := string(board.OwnerOf(backend.A2))
	a3 := string(board.OwnerOf(backend.A3))
	b1 := string(board.OwnerOf(backend.B1))
	b2 := string(board.OwnerOf(backend.B2))
	b3 := string(board.OwnerOf(backend.B3))
	c1 := string(board.OwnerOf(backend.C1))
	c2 := string(board.OwnerOf(backend.C2))
	c3 := string(board.OwnerOf(backend.C3))
	res.A1 = &a1
	res.A2 = &a2
	res.A3 = &a3
	res.B1 = &b1
	res.B2 = &b2
	res.B3 = &b3
	res.C1 = &c1
	res.C2 = &c2
	res.C3 = &c3
	return
}
