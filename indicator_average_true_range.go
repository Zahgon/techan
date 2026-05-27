package techan

import "github.com/sdcoffey/big"

type averageTrueRangeIndicator struct {
	series *TimeSeries
	window int
}

// NewAverageTrueRangeIndicator returns a base indicator that calculates the average true range of the
// underlying over a window
// https://www.investopedia.com/terms/a/atr.asp
func NewAverageTrueRangeIndicator(series *TimeSeries, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (atr averageTrueRangeIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
