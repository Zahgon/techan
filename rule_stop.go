package techan

import "github.com/sdcoffey/big"

type stopLossRule struct {
	Indicator
	tolerance big.Decimal
}

// NewStopLossRule returns a new rule that is satisfied when the given loss tolerance (a percentage) is met or exceeded.
// Loss tolerance should be a value between -1 and 1.
func NewStopLossRule(series *TimeSeries, lossTolerance float64) Rule {
	_ = "STUB: not implemented"
	return *new(Rule)
}

func (slr stopLossRule) IsSatisfied(index int, record *TradingRecord) bool {
	_ = "STUB: not implemented"
	return false
}
