package techan

import "github.com/sdcoffey/big"

type emaIndicator struct {
	indicator   Indicator
	window      int
	alpha       big.Decimal
	resultCache resultCache
}

// NewEMAIndicator returns a derivative indicator which returns the average of the current and preceding values in
// the given windowSize, with values closer to current index given more weight. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/e/ema.asp
func NewEMAIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (ema *emaIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

func (ema emaIndicator) cache() resultCache { _ = "STUB: not implemented"; return *new(resultCache) }

func (ema *emaIndicator) setCache(newCache resultCache) { _ = "STUB: not implemented"; return }

func (ema emaIndicator) windowSize() int { _ = "STUB: not implemented"; return 0 }
