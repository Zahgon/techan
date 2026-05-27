package techan

import "github.com/sdcoffey/big"

type resultCache []*big.Decimal

type cachedIndicator interface {
	Indicator
	cache() resultCache
	setCache(cache resultCache)
	windowSize() int
}

func cacheResult(indicator cachedIndicator, index int, val big.Decimal) {
	_ = "STUB: not implemented"
	return
}

func expandResultCache(indicator cachedIndicator, newSize int) { _ = "STUB: not implemented"; return }

func returnIfCached(indicator cachedIndicator, index int, firstValueFallback func(int) big.Decimal) *big.Decimal {
	_ = "STUB: not implemented"
	return nil
}
