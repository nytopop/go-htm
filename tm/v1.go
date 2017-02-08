package tm

import "github.com/nytopop/gohtm/cells"

// V1Params contains parameters for initialization of an Extended
// TemporalMemory region.
type V1Params struct {
	NumColumns  int // input space dimensions
	CellsPerCol int // cells per column
	SegsPerCell int
	SynsPerSeg  int
	MaxNewSyns  int // max new synapses per segment per iteration
}

// NewV1Params returns a default V1Params.
func NewV1Params() V1Params {
	return V1Params{
		NumColumns:  2048,
		CellsPerCol: 32,
		SegsPerCell: 16,
		SynsPerSeg:  16,
		MaxNewSyns:  20,
	}
}

// V1 is a TemporalMemory implementation with support
// for basal and apical dendrites connected to other regions.
type V1 struct {
	P    V1Params
	Cons cells.Cells

	PrevActiveCells []bool
	PrevWinnerCells []bool
	ActiveCells     []bool
	WinnerCells     []bool
	iteration       int
}

// NewV1 initializes a new TemporalMemory region with the provided V1Params.
func NewV1(p V1Params) *V1 {
	cp := cells.NewV1Params()
	cp.NumColumns = p.NumColumns
	cp.CellsPerCol = p.CellsPerCol
	cp.SegsPerCell = p.SegsPerCell
	cp.SynsPerSeg = p.SynsPerSeg

	return &V1{
		P:               p,
		Cons:            cells.NewV1(cp),
		PrevActiveCells: make([]bool, 0, p.NumColumns*p.CellsPerCol),
		PrevWinnerCells: make([]bool, 0, p.NumColumns*p.CellsPerCol),
		ActiveCells:     make([]bool, 0, p.NumColumns*p.CellsPerCol),
		WinnerCells:     make([]bool, 0, p.NumColumns*p.CellsPerCol),
		iteration:       0,
	}
}

// Compute iterates the TemporalMemory algorithm with the
// provided vector of active columns from a SpatialPooler.
func (e *V1) Compute(active []bool, learn bool) {
	e.iteration += 1
	e.activateCells(active, learn)
	e.activateDendrites(learn)
}

// Calculate the active cells using active columns and dendrite segments.
// Grow and reinforce synapses.
func (e *V1) activateCells(active []bool, learn bool) {
	/*
		for each column
		  if col active and has active dendrite segments
		    call activatePredictedColumn
		  if col active and doesn't have active dendrite segments
		    call burstColumn
		  if col inactive and has matching dendrite segments
		    call punishpredictedcolumn
	*/

	// Allocate memory for active / winner cells
	e.PrevActiveCells = e.ActiveCells
	e.PrevWinnerCells = e.WinnerCells
	e.ActiveCells = make([]bool, e.P.NumColumns*e.P.CellsPerCol)
	e.WinnerCells = make([]bool, e.P.NumColumns*e.P.CellsPerCol)

	// TODO : run a count on how big activecells and winnercells
	// 		  actually get, probably don't need full (cols * cellspercol)
	for i := range active {
		if active[i] {
			// if has active dendrite segments, activatePredictedCol
			// if not, burst
			syns := e.Cons.ActiveSegsForCol(i)
			switch {
			case syns > 0:
				cellsToAdd := e.activatePredictedColumn(i, learn)
				for _, c := range cellsToAdd {
					e.ActiveCells[c] = true
					e.WinnerCells[c] = true
				}
			case syns == 0:
				cellsToAdd, winnerCell := e.burstColumn(i, learn)
				for _, c := range cellsToAdd {
					e.ActiveCells[c] = true
				}
				e.WinnerCells[winnerCell] = true
			}
		} else {
			// if has matching dendrite segments,
			// punishPredictedColumn
			if learn {
				if e.Cons.MatchingSegsForCol(i) > 0 {
					e.punishPredictedColumn(i)
				}
			}
		}
	}
}

func (e *V1) activatePredictedColumn(col int, learn bool) []int {
	/*
		for each cell in col that has an active distal dendrite segment
		  mark cell as active cell
		  mark cell as winner cell
		  if learning enabled
		    strengthen active synapses
			weaken inactive synapses
			grow synapses to previous winner cells
	*/
	cellsToAdd := make([]int, 0, e.P.CellsPerCol) // TODO sizing ???

	for _, i := range e.Cons.CellsForCol(col) {
		if e.Cons.ActiveSegsForCell(i) > 0 {
			cellsToAdd = append(cellsToAdd, i)

			if learn {
				e.Cons.AdaptSynapses(i, e.PrevActiveCells)
				e.Cons.GrowSynapses(i, e.PrevWinnerCells)
				// grow synapses to prev winner cells IF below max seg count
			}
		}
	}

	return cellsToAdd
}

func (e *V1) burstColumn(col int, learn bool) ([]int, int) {
	cellsToAdd := make([]int, 0, e.P.NumColumns) // TODO sizing ???
	winnerCell := 0
	return cellsToAdd, winnerCell
}

func (e *V1) punishPredictedColumn(col int) {
}

func (e *V1) activateDendrites(learn bool) {
	/*
		for each
	*/
	if learn {
		e.Cons.StartNewIteration()
	}

	// Compute active & matching dendrite segments
	e.Cons.Clear()
	e.Cons.ComputeActivity(e.ActiveCells)
}

// Reset clears temporary data so sequences are not learned between
// the current and next time step.
func (e *V1) Reset() {
	e.PrevActiveCells = make([]bool, e.P.NumColumns*e.P.CellsPerCol)
	e.PrevWinnerCells = make([]bool, e.P.NumColumns*e.P.CellsPerCol)
	e.ActiveCells = make([]bool, e.P.NumColumns*e.P.CellsPerCol)
	e.WinnerCells = make([]bool, e.P.NumColumns*e.P.CellsPerCol)
}
