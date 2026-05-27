package techan

import "github.com/sdcoffey/big"

// NewMaximumValueIndicator returns a derivative Indicator which returns the maximum value
// present in a given window. Use a window value of -1 to include all values in the
// underlying indicator.
func NewMaximumValueIndicator(ind Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

type maximumValueIndicator struct {
	indicator Indicator
	window    int
}

func (mvi maximumValueIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
