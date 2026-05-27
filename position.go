package techan

import "github.com/sdcoffey/big"

// Position is a pair of two Order objects
type Position struct {
	orders [2]*Order
}

// NewPosition returns a new Position with the passed-in order as the open order
func NewPosition(openOrder Order) (t *Position) { _ = "STUB: not implemented"; return nil }

// Enter sets the open order to the order passed in
func (p *Position) Enter(order Order) { _ = "STUB: not implemented"; return }

// Exit sets the exit order to the order passed in
func (p *Position) Exit(order Order) { _ = "STUB: not implemented"; return }

// IsLong returns true if the entrance order is a buy order
func (p *Position) IsLong() bool { _ = "STUB: not implemented"; return false }

// IsShort returns true if the entrance order is a sell order
func (p *Position) IsShort() bool { _ = "STUB: not implemented"; return false }

// IsOpen returns true if there is an entrance order but no exit order
func (p *Position) IsOpen() bool { _ = "STUB: not implemented"; return false }

// IsClosed returns true of there are both entrance and exit orders
func (p *Position) IsClosed() bool { _ = "STUB: not implemented"; return false }

// IsNew returns true if there is neither an entrance or exit order
func (p *Position) IsNew() bool { _ = "STUB: not implemented"; return false }

// EntranceOrder returns the entrance order of this position
func (p *Position) EntranceOrder() *Order {
	_ = "STUB: not implemented"

	// ExitOrder returns the exit order of this position
	return nil
}

func (p *Position) ExitOrder() *Order {
	_ = "STUB: not implemented"

	// CostBasis returns the price to enter this order
	return nil
}

func (p *Position) CostBasis() big.Decimal { _ = "STUB: not implemented"; return *new(big.Decimal) }

// ExitValue returns the value accrued by closing the position
func (p *Position) ExitValue() big.Decimal { _ = "STUB: not implemented"; return *new(big.Decimal) }
