package techan

//lint:file-ignore S1038 prefer Fprintln

import (
	"io"
	"time"
)

// Analysis is an interface that describes a methodology for taking a TradingRecord as input,
// and giving back some float value that describes it's performance with respect to that methodology.
type Analysis interface {
	Analyze(*TradingRecord) float64
}

// TotalProfitAnalysis analyzes the trading record for total profit.
type TotalProfitAnalysis struct{}

// Analyze analyzes the trading record for total profit.
func (tps TotalProfitAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// PercentGainAnalysis analyzes the trading record for the percentage profit gained relative to start
type PercentGainAnalysis struct{}

// Analyze analyzes the trading record for the percentage profit gained relative to start
func (pga PercentGainAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// NumTradesAnalysis analyzes the trading record for the number of trades executed
type NumTradesAnalysis string

// Analyze analyzes the trading record for the number of trades executed
func (nta NumTradesAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// LogTradesAnalysis is a wrapper around an io.Writer, which logs every trade executed to that writer
type LogTradesAnalysis struct {
	io.Writer
}

// Analyze logs trades to provided io.Writer
func (lta LogTradesAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// PeriodProfitAnalysis analyzes the trading record for the average profit based on the time period provided.
// i.e., if the trading record spans a year of trading, and PeriodProfitAnalysis wraps one month, Analyze will return
// the total profit for the whole time period divided by 12.
type PeriodProfitAnalysis struct {
	Period time.Duration
}

// Analyze returns the average profit for the trading record based on the given duration
func (ppa PeriodProfitAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// ProfitableTradesAnalysis analyzes the trading record for the number of profitable trades
type ProfitableTradesAnalysis struct{}

// Analyze returns the number of profitable trades in a trading record
func (pta ProfitableTradesAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// AverageProfitAnalysis returns the average profit for the trading record. Average profit is represented as the total
// profit divided by the number of trades executed.
type AverageProfitAnalysis struct{}

// Analyze returns the average profit of the trading record
func (apa AverageProfitAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}

// BuyAndHoldAnalysis returns the profit based on a hypothetical where a purchase order was made on the first period available
// and held until the date on the last trade of the trading record. It's useful for comparing the performance of your strategy
// against a simple long position.
type BuyAndHoldAnalysis struct {
	TimeSeries    *TimeSeries
	StartingMoney float64
}

// Analyze returns the profit based on a simple buy and hold strategy
func (baha BuyAndHoldAnalysis) Analyze(record *TradingRecord) float64 {
	_ = "STUB: not implemented"
	return 0
}
