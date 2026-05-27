package techan

import "github.com/sdcoffey/big"

type volumeIndicator struct {
	*TimeSeries
}

// NewVolumeIndicator returns an indicator which returns the volume of a candle for a given index
func NewVolumeIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (vi volumeIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type closePriceIndicator struct {
	*TimeSeries
}

// NewClosePriceIndicator returns an Indicator which returns the close price of a candle for a given index
func NewClosePriceIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (cpi closePriceIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type highPriceIndicator struct {
	*TimeSeries
}

// NewHighPriceIndicator returns an Indicator which returns the high price of a candle for a given index
func NewHighPriceIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (hpi highPriceIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type lowPriceIndicator struct {
	*TimeSeries
}

// NewLowPriceIndicator returns an Indicator which returns the low price of a candle for a given index
func NewLowPriceIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (lpi lowPriceIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type openPriceIndicator struct {
	*TimeSeries
}

// NewOpenPriceIndicator returns an Indicator which returns the open price of a candle for a given index
func NewOpenPriceIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (opi openPriceIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}

type typicalPriceIndicator struct {
	*TimeSeries
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewTypicalPriceIndicator(series *TimeSeries) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (tpi typicalPriceIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
