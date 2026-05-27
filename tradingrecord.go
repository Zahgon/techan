package techan

// TradingRecord is an object describing a series of trades made and a current position
type TradingRecord struct {
	Trades          []*Position
	currentPosition *Position
}

// NewTradingRecord returns a new TradingRecord
func NewTradingRecord() (t *TradingRecord) { _ = "STUB: not implemented"; return nil }

// CurrentPosition returns the current position in this record
func (tr *TradingRecord) CurrentPosition() *Position { _ = "STUB: not implemented"; return nil }

// LastTrade returns the last trade executed in this record
func (tr *TradingRecord) LastTrade() *Position { _ = "STUB: not implemented"; return nil }

// Operate takes an order and adds it to the current TradingRecord. It will only add the order if:
// - The current position is open and the passed order was executed after the entrance order
// - The current position is new and the passed order was executed after the last exit order
func (tr *TradingRecord) Operate(order Order) { _ = "STUB: not implemented"; return }
