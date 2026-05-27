package techan

import (
	"github.com/sdcoffey/big"
)

type kIndicator struct {
	closePrice Indicator
	minValue   Indicator
	maxValue   Indicator
	window     int
}

// NewFastStochasticIndicator returns a derivative Indicator which returns the fast stochastic indicator (%K) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewFastStochasticIndicator(series *TimeSeries, timeframe int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (k kIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type dIndicator struct {
	k      Indicator
	window int
}

// NewSlowStochasticIndicator returns a derivative Indicator which returns the slow stochastic indicator (%D) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewSlowStochasticIndicator(k Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (d dIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
