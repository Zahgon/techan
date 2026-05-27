package techan

import "github.com/sdcoffey/big"

// NewVarianceIndicator provides a way to find the variance in a base indicator, where variances is the sum of squared
// deviations from the mean at any given index in the time series.
func NewVarianceIndicator(ind Indicator) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

type varianceIndicator struct {
	Indicator Indicator
}

// Calculate returns the Variance for this indicator at the given index
func (vi varianceIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
