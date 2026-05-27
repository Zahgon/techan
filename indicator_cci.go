package techan

import "github.com/sdcoffey/big"

type commidityChannelIndexIndicator struct {
	series *TimeSeries
	window int
}

// NewCCIIndicator Returns a new Commodity Channel Index Indicator
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:commodity_channel_index_cci
func NewCCIIndicator(ts *TimeSeries, window int) Indicator {
	_ = "STUB: not implemented"
	return *new(Indicator)
}

func (ccii commidityChannelIndexIndicator) Calculate(index int) big.Decimal {
	_ = "STUB: not implemented"
	return *new(big.Decimal)
}
