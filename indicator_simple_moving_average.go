package techan

import "github.com/sdcoffey/big"

type smaIndicator struct {
	indicator Indicator
	window    int
}

// NewSimpleMovingAverage returns a derivative Indicator which returns the average of the current value and preceding
// values in the given windowSize.
func NewSimpleMovingAverage(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (sma smaIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
