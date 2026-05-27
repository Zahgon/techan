package techan

import "github.com/sdcoffey/big"

type trendLineIndicator struct {
	indicator Indicator
	window    int
}

// NewTrendlineIndicator returns an indicator whose output is the slope of the trend
// line given by the values in the window.
func NewTrendlineIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (tli trendLineIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

func sumX(decimals []big.Decimal) (s big.Decimal) {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

func sumY(decimals []big.Decimal) (b big.Decimal) {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

func sumXy(decimals []big.Decimal) (b big.Decimal) {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

func sumX2(decimals []big.Decimal) big.Decimal { _ = "STUB: not implemented"; return *new(big.Decimal) }
