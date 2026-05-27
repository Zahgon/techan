package techan

import "github.com/sdcoffey/big"

type meanDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

// NewMeanDeviationIndicator returns a derivative Indicator which returns the mean deviation of a base indicator
// in a given window. Mean deviation is an average of all values on the base indicator from the mean of that indicator.
func NewMeanDeviationIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (mdi meanDeviationIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
