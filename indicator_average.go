package techan

import "github.com/sdcoffey/big"

type averageIndicator struct {
	Indicator
	window int
}

// NewAverageGainsIndicator Returns a new average gains indicator, which returns the average gains
// in the given window based on the given indicator.
func NewAverageGainsIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

// NewAverageLossesIndicator Returns a new average losses indicator, which returns the average losses
// in the given window based on the given indicator.
func NewAverageLossesIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (ai averageIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
