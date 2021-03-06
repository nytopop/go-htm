package enc

// TODO : does not implement Encoder

// ScalarParams represents a parameter set for a Scalar Encoder.
type ScalarParams struct {
	Buckets  int
	Min, Max int
	Active   int
	Wrap     bool
}

// NewScalarParams returns a default param set.
func NewScalarParams() ScalarParams {
	return ScalarParams{
		Buckets: 66,
		Min:     0,
		Max:     65,
		Active:  10,
		Wrap:    false,
	}
}

// Scalar is a linearly derived scalar encoder.
type Scalar struct {
	P     ScalarParams
	Bits  int
	Range int
}

// NewScalar returns a Scalar encoder initialiazed with the provided
// ScalarParams.
func NewScalar(p ScalarParams) *Scalar {
	return &Scalar{
		P:     p,
		Bits:  p.Buckets + p.Active - 1,
		Range: p.Max - p.Min, // should be Absolute value
	}
}

// Encode asdf
func (s *Scalar) Encode(d interface{}) []bool {
	out := make([]bool, s.Bits)
	i := s.P.Buckets * (d.(int) - s.P.Min) / s.Range
	for j := 0; j < s.P.Active; j++ {
		if (i + j) < (len(out) - 1) {
			out[i+j] = true
		} else {
			if s.P.Wrap {
				// TODO wraparound
			}
		}
	}
	return out
}

// Decode asdf
func (s *Scalar) Decode(sv []bool) interface{} {
	for i, v := range sv {
		if v {
			return (i+1)*s.Range/s.P.Buckets - s.P.Min
		}
	}
	return 0
}
