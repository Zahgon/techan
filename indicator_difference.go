package techan

import "github.com/sdcoffey/big"

type differenceIndicator struct {
	minuend    Indicator
	subtrahend Indicator
}

// NewDifferenceIndicator returns an indicator which returns the difference between one indicator (minuend) and a second
// indicator (subtrahend).
func NewDifferenceIndicator(minuend, subtrahend Indicator) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (di differenceIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
