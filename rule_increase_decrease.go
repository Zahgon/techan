package techan

// IncreaseRule is satisfied when the given Indicator at the given index is greater than the value at the previous
// index.
type IncreaseRule struct {
	Indicator
}

// IsSatisfied returns true when the given Indicator at the given index is greater than the value at the previous
// index.
func (ir IncreaseRule) IsSatisfied(index int, record *TradingRecord) bool {
	_ = "STUB: not implemented"
	return false
}

// DecreaseRule is satisfied when the given Indicator at the given index is less than the value at the previous
// index.
type DecreaseRule struct {
	Indicator
}

// IsSatisfied returns true when the given Indicator at the given index is less than the value at the previous
// index.
func (dr DecreaseRule) IsSatisfied(index int, record *TradingRecord) bool {
	_ = "STUB: not implemented"
	return false
}
