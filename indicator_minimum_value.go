package techan

import "github.com/sdcoffey/big"

// NewMinimumValueIndicator returns a derivative Indicator which returns the minimum value
// present in a given window. Use a window value of -1 to include all values in the
// underlying indicator.
func NewMinimumValueIndicator(ind Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

type minimumValueIndicator struct {
	indicator Indicator
	window    int
}

func (mvi minimumValueIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
