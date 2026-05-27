package techan

import "github.com/sdcoffey/big"

type modifiedMovingAverageIndicator struct {
	indicator   Indicator
	window      int
	resultCache resultCache
}

// NewMMAIndicator returns a derivative indciator which returns the modified moving average of the underlying
// indictator. An in-depth explanation can be found here:
// https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average
func NewMMAIndicator(indicator Indicator, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (mma *modifiedMovingAverageIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

func (mma modifiedMovingAverageIndicator) cache() resultCache {
	_ = "STUB: not implemented"
	return *new(resultCache)
}

func (mma *modifiedMovingAverageIndicator) setCache(cache resultCache) {
	_ = "STUB: not implemented"
	return
}

func (mma modifiedMovingAverageIndicator) windowSize() int { _ = "STUB: not implemented"; return 0 }
