package techan

import (
	"github.com/sdcoffey/big"
)

type keltnerChannelIndicator struct {
	ema    Indicator
	atr    Indicator
	mul    big.Decimal
	window int
}

func NewKeltnerChannelUpperIndicator(series *TimeSeries, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func NewKeltnerChannelLowerIndicator(series *TimeSeries, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (kci keltnerChannelIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
