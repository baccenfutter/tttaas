package tictactoe

// allSquares is a slice containing all squares in descending order.
var allSquares = []Square{
	A1, A2, A3,
	B1, B2, B3,
	C1, C2, C3,
}

// max will return the largest of the given integers.
func max(v ...int) int {
	out := 0
	for i := range v {
		if v[i] > out {
			out = v[i]
		}
	}
	return out
}

// min will returns the smallest of the given integers
func min(v ...int) int {
	out := 0
	if len(v) >= 1 {
		out = v[0]
	}
	for i := range v {
		if v[i] < out {
			out = v[i]
		}
	}
	return out
}

// playerInList returns true if the given Player is contained within the given []Player.
func playerInList(p Player, ls []Player) bool {
	for _, player := range ls {
		if p == player {
			return true
		}
	}
	return false
}

// vacantSquares returns a []Square{} containing all vacant squares on this board.
func (b *Board) vacantSquares() []Square {
	occupiedSquares := append(b.Player1, b.Player2...)
	out := []Square{}
	for _, sq := range allSquares {
		found := false
		for _, occupiedSquare := range occupiedSquares {
			if sq == occupiedSquare {
				found = true
				break
			}
		}
		if !found {
			out = append(out, Square(sq))
		}
	}
	return out
}

// copy returns a copy of this board.
// This function is mainly used to create immutable clones for evaluating moves.
func (b *Board) copy() *Board {
	newBoard := &Board{
		Player1: make([]Square, len(b.Player1)),
		Player2: make([]Square, len(b.Player2)),
		result:  b.result,
	}
	for i := range b.Player1 {
		newBoard.Player1[i] = b.Player1[i]
	}
	for i := range b.Player2 {
		newBoard.Player2[i] = b.Player2[i]
	}
	return newBoard
}

// boardEvaluatoar is a helper for evaluating a board position.
type boardEvaluator struct {
	Rows map[Row]map[Square]bool
	Cols map[Col]map[Square]bool
	Diag [2]map[Square]bool
}

// newBoardEvaluator returns an initialized *boardEvaluator{} based on the given []Square.
func newBoardEvaluator(squares []Square) *boardEvaluator {
	b := &boardEvaluator{
		Rows: map[Row]map[Square]bool{
			RowA: map[Square]bool{},
			RowB: map[Square]bool{},
			RowC: map[Square]bool{},
		},
		Cols: map[Col]map[Square]bool{
			Col1: map[Square]bool{},
			Col2: map[Square]bool{},
			Col3: map[Square]bool{},
		},
		Diag: [2]map[Square]bool{
			map[Square]bool{},
			map[Square]bool{},
		},
	}
	for _, sq := range squares {
		b.Rows[sq.Row()][sq] = true
		b.Cols[sq.Col()][sq] = true
		if sq == A1 || sq == B2 || sq == C3 {
			b.Diag[0][sq] = true
		}
		if sq == A3 || sq == B2 || sq == C1 {
			b.Diag[1][sq] = true
		}
	}
	return b
}

// isWin returns true if the loaded position is a winning position.
func (eval *boardEvaluator) isWin() bool {
	for _, row := range eval.Rows {
		if len(row) == 3 {
			return true
		}
	}
	for _, col := range eval.Cols {
		if len(col) == 3 {
			return true
		}
	}
	for _, diag := range eval.Diag {
		if len(diag) == 3 {
			return true
		}
	}
	return false
}

// possibleMoves returns a []*Board{} with copies of this board for each possible move
// on this board.
func (b *Board) possibleMoves() []*Board {
	out := []*Board{}
	for _, sq := range b.vacantSquares() {
		newBoard := b.copy()
		newBoard.Set(sq)
		out = append(out, newBoard)
	}
	return out
}

// valuate_static returns which Player would be winning if this position was final.
func (b *Board) evaluate_static() Player {
	if newBoardEvaluator(b.Player1).isWin() {
		return Player1
	} else if newBoardEvaluator(b.Player2).isWin() {
		return Player2
	}
	return NoPlayer
}

// score implements the minimax algorithm.
func (b *Board) score() Player {
	static_eval := b.evaluate_static()
	if static_eval != NoPlayer {
		return static_eval
	}

	if len(b.vacantSquares()) == 0 {
		return NoPlayer
	}

	branch_evals := func() []Player {
		out := []Player{}
		for _, branch := range b.possibleMoves() {
			out = append(out, branch.score())
		}
		return out
	}()

	switch b.NextPlayer() {
	case Player1:
		if playerInList(Player1, branch_evals) {
			return Player1
		} else if playerInList(NoPlayer, branch_evals) {
			return NoPlayer
		} else {
			return Player2
		}
	case Player2:
		if playerInList(Player2, branch_evals) {
			return Player2
		} else if playerInList(NoPlayer, branch_evals) {
			return NoPlayer
		} else {
			return Player1
		}
	}

	return NoPlayer
}

// MakeOptimalMove makes an optimal move on the given board.
func MakeOptimalMove(b *Board) {
	aiPlayer := b.NextPlayer()
	branch_evals := func() []Player {
		out := []Player{}
		for _, branch := range b.possibleMoves() {
			out = append(out, branch.score())
		}
		return out
	}()

	var choice Player
	switch aiPlayer {
	case Player1:
		if playerInList(Player1, branch_evals) {
			choice = Player1
		} else if playerInList(NoPlayer, branch_evals) {
			choice = NoPlayer
		} else {
			choice = Player2
		}
	case Player2:
		if playerInList(Player2, branch_evals) {
			choice = Player2
		} else if playerInList(NoPlayer, branch_evals) {
			choice = NoPlayer
		} else {
			choice = Player1
		}
	}
	for i, sq := range b.possibleMoves() {
		if sq.score() == choice {
			b.Set(b.vacantSquares()[i])
			return
		}
	}
}
