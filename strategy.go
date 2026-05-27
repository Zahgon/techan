package techan

// Strategy is an interface that describes desired entry and exit trading behavior
type Strategy interface {
	ShouldEnter(index int, record *TradingRecord) bool
	ShouldExit(index int, record *TradingRecord) bool
}

// RuleStrategy is a strategy based on rules and an unstable period. The two rules determine whether a position should
// be created or closed, and the unstable period is an index before no positions should be created or exited
type RuleStrategy struct {
	EntryRule      Rule
	ExitRule       Rule
	UnstablePeriod int
}

// ShouldEnter will return true when the index is less than the unstable period and the entry rule is satisfied
func (rs RuleStrategy) ShouldEnter(index int, record *TradingRecord) bool {
	_ = "STUB: not implemented"
	return false
}

// ShouldExit will return true when the index is less than the unstable period and the exit rule is satisfied
func (rs RuleStrategy) ShouldExit(index int, record *TradingRecord) bool {
	_ = "STUB: not implemented"
	return false
}
