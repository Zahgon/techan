package techan

import "github.com/sdcoffey/big"

type windowedStandardDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

// NewWindowedStandardDeviationIndicator returns a indicator which calculates the standard deviation of the underlying
// indicator over a window
func NewWindowedStandardDeviationIndicator(ind Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (sdi windowedStandardDeviationIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
