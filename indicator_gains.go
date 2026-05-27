package techan

import "github.com/sdcoffey/big"

type gainLossIndicator struct {
	Indicator
	coefficient big.Decimal
}

// NewGainIndicator returns a derivative indicator that returns the gains in the underlying indicator in the last bar,
// if any. If the delta is negative, zero is returned
func NewGainIndicator(indicator Indicator) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

// NewLossIndicator returns a derivative indicator that returns the losses in the underlying indicator in the last bar,
// if any. If the delta is positive, zero is returned
func NewLossIndicator(indicator Indicator) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (gli gainLossIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type cumulativeIndicator struct {
	Indicator
	window int
	mult   big.Decimal
}

// NewCumulativeGainsIndicator returns a derivative indicator which returns all gains made in a base indicator for a given
// window.
func NewCumulativeGainsIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

// NewCumulativeLossesIndicator returns a derivative indicator which returns all losses in a base indicator for a given
// window.
func NewCumulativeLossesIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (ci cumulativeIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type percentChangeIndicator struct {
	Indicator
}

// NewPercentChangeIndicator returns a derivative indicator which returns the percent change (positive or negative)
// made in a base indicator up until the given indicator
func NewPercentChangeIndicator(indicator Indicator) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (pgi percentChangeIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
