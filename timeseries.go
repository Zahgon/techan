package techan

// TimeSeries represents an array of candles
type TimeSeries struct {
	Candles []*Candle
}

// NewTimeSeries returns a new, empty, TimeSeries
func NewTimeSeries() (t *TimeSeries) { _ = "STUB: not implemented"; return nil }

// AddCandle adds the given candle to this TimeSeries if it is not nil and after the last candle in this timeseries.
// If the candle is added, AddCandle will return true, otherwise it will return false.
func (ts *TimeSeries) AddCandle(candle *Candle) bool { _ = "STUB: not implemented"; return false }

// LastCandle will return the lastCandle in this series, or nil if this series is empty
func (ts *TimeSeries) LastCandle() *Candle { _ = "STUB: not implemented"; return nil }

// LastIndex will return the index of the last candle in this series
func (ts *TimeSeries) LastIndex() int { _ = "STUB: not implemented"; return 0 }
