package techan

import "github.com/sdcoffey/big"

type fixedIndicator []float64

// NewFixedIndicator returns an indicator with a fixed set of values that are returned when an index is passed in
func NewFixedIndicator(vals ...float64) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (fi fixedIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
