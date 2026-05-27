package techan

import "github.com/sdcoffey/big"

type trueRangeIndicator struct {
	series *TimeSeries
}

// NewTrueRangeIndicator returns a base indicator
// which calculates the true range at the current point in time for a series
// https://www.investopedia.com/terms/a/atr.asp
func NewTrueRangeIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (tri trueRangeIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
