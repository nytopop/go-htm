package main

/* Spatial Pooling
1. Start with input of fixed n bits.
2. Assign fixed number of columns to region receiving the input. Each column has associated dendrite segment. Each dendrite segment has a set of potential synapses representing subset of input bits. Each potential synapse has a permanence value; based on that value, some potential synapses will be connected to the dendrite segment.
3. For any input, determine how many connected synapses on a column are connected to active input bits.
4. Number of active synapses is multiplied by a boosting factor; dynamically determined by how often a column is active relative to neighbors.
5. A fixed n of columns within the inhibition radius with highest activations after boosting become active and disable rest of columns within the radius. The inhibition radius is dynamically determined by the spread of input bits. We should now have a sparse set of active columns.
6. For each active column, we adjust the permanence values of all potential synapses. The permanence of synapses aligned with active input bits is increased. The changes may change some synapses from connected<->disconnected.
*/

/*
Need to convert input vector into output vector of dif dimensions
Need to normalize sparsity
Maintain overlap properties of input data
*/

type ProximalSynapse struct {
}

type DistalSynapse struct {
}

type Dendrite struct {
}

type Column struct {
	d []Dendrite
}

type SpatialPooler struct {
	// Params
	n           int     // number of columns
	h           int     // column height
	connPercent float64 // percent to connect on init

	// State
	columns   []Column // columns
	potential Vector
}

func NewSpatialPooler() SpatialPooler {
	sp := SpatialPooler{}

	return sp
}

func (sp SpatialPooler) Compute(in SparseVector) SparseVector {
	// phase 1
	// phase 2
	// phase 3
	return SparseVector{}
}

// Phase 1 : Overlap
func (sp SpatialPooler) Overlap() {
}

// Phase 2 : Inhibition
func (sp SpatialPooler) Inhibit() {
}

// Phase 3 : Learning
func (sp SpatialPooler) Learn() {
}