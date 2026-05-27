package techan

import "github.com/sdcoffey/big"

type relativeVigorIndexIndicator struct {
	numerator   Indicator
	denominator Indicator
}

// NewRelativeVigorIndexIndicator returns an Indicator which returns the index of the relative vigor of the prices of
// a sercurity. Relative Vigor Index is simply the difference of the previous four days' close and open prices divided
// by the difference between the previous four days high and low prices. A more in-depth explanation of relative vigor
// index can be found here: https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/relative-vigor-index
func NewRelativeVigorIndexIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (rvii relativeVigorIndexIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type relativeVigorIndexSignalLine struct {
	relativeVigorIndex Indicator
}

// NewRelativeVigorSignalLine returns an Indicator intended to be used in conjunction with Relative vigor index, which
// returns the average value of the last 4 indices of the RVI indicator.
func NewRelativeVigorSignalLine(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (rvsn relativeVigorIndexSignalLine) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
