package gohtm

import "fmt"

/* Temporal Memory */

// TemporalParams contains parameters for initialization of a
// TemporalMemory region.
type TemporalParams struct {
	NumColumns          int     // input space dimensions
	NumCells            int     // cells per column
	ActivationThreshold int     // # of active synapses for 'active'
	MinThreshold        int     // # of potential synapses for learning
	InitPermanence      float64 // initial permanence of new synapses
	SynPermConnected    float64 // connection threshold
	SegPerCell          int
	SynPerSeg           int
}

// NewTemporalParams returns a default TemporalParams.
func NewTemporalParams() TemporalParams {
	return TemporalParams{
		NumColumns:          128,
		NumCells:            4,
		ActivationThreshold: 12,
		MinThreshold:        10,
		InitPermanence:      0.21,
		SynPermConnected:    0.5,
		SegPerCell:          64,
		SynPerSeg:           64,
	}
}

// TemporalMemory is a sequence learning and prediction algorithm.
type TemporalMemory struct {
	// state
	cons Connections

	// params
	numColumns          int
	numCells            int
	activationThreshold int
	minThreshold        int
	initPermanence      float64
	synPermConnected    float64
	segPerCell          int
	synPerSeg           int
}

// NewTemporalMemory initializes a new TemporalMemory region with
// the provided TemporalParams.
func NewTemporalMemory(p TemporalParams) TemporalMemory {
	tm := TemporalMemory{
		numColumns:          p.NumColumns,
		numCells:            p.NumCells,
		activationThreshold: p.ActivationThreshold,
		minThreshold:        p.MinThreshold,
		initPermanence:      p.InitPermanence,
		synPermConnected:    p.SynPermConnected,
		segPerCell:          p.SegPerCell,
		synPerSeg:           p.SynPerSeg,
	}

	totalCells := tm.numColumns * tm.numCells
	fmt.Println(totalCells, "total cells")
	tm.cons = NewConnections(totalCells, tm.segPerCell, tm.synPerSeg)
	fmt.Println(tm.cons)

	/*
		tm.cols = make([]TMColumn, p.numColumns)
		for i, _ := range tm.cols {
			// i is column index
			tm.cols[i] = make(TMColumn, p.numCells)

			for j, _ := range tm.cols[i] {
				// j is cell index

				tm.cols[i][j].seg = Segment{
					dsyns: make([]DistalSynapse, 128),
				}
			}
		}
	*/

	return tm
}

// Compute iterates the TemporalMemory algorithm with the provided
// vector of active columns from a SpatialPooler.
func (tm *TemporalMemory) Compute(active SparseBinaryVector) SparseBinaryVector {
	return SparseBinaryVector{}
}
