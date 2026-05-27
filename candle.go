package techan

import (
	"github.com/sdcoffey/big"
)

// Candle represents basic market information for a security over a given time period
type Candle struct {
	Period     TimePeriod
	OpenPrice  big.Decimal
	ClosePrice big.Decimal
	MaxPrice   big.Decimal
	MinPrice   big.Decimal
	Volume     big.Decimal
	TradeCount uint
}

// NewCandle returns a new *Candle for a given time period
func NewCandle(period TimePeriod) (c *Candle) { _ = "STUB: not implemented"; return nil }

// AddTrade adds a trade to this candle. It will determine if the current price is higher or lower than the min or max
// price and increment the tradecount.
func (c *Candle) AddTrade(tradeAmount, tradePrice big.Decimal) { _ = "STUB: not implemented"; return }

func (c *Candle) String() string { _ = "STUB: not implemented"; return "" }
