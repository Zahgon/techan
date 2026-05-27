package techan

import (
	"github.com/sdcoffey/big"
)

type aroonIndicator struct {
	indicator Indicator
	window    int
	direction big.Decimal
	lowIndex  int
}

func (ai *aroonIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

func (ai aroonIndicator) findLowIndex(index int) int { _ = "STUB: not implemented"; return 0 }

// NewAroonUpIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the highest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a HighPriceIndicator or a derivative thereof
func NewAroonUpIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

// NewAroonDownIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the lowest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a LowPriceIndicator or a derivative thereof
func NewAroonDownIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}
