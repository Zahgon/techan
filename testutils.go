package techan

import (
	"testing"

	"github.com/sdcoffey/big"
)

var candleIndex int
var mockedTimeSeries = mockTimeSeriesFl(
	64.75, 63.79, 63.73,
	63.73, 63.55, 63.19,
	63.91, 63.85, 62.95,
	63.37, 61.33, 61.51)

func randomTimeSeries(size int) *TimeSeries { _ = "STUB: not implemented"; return nil }

func mockTimeSeriesOCHL(values ...[]float64) *TimeSeries { _ = "STUB: not implemented"; return nil }

func mockTimeSeries(values ...string) *TimeSeries { _ = "STUB: not implemented"; return nil }

func mockTimeSeriesFl(values ...float64) *TimeSeries { _ = "STUB: not implemented"; return nil }

func decimalEquals(t *testing.T, expected float64, actual big.Decimal) {
	_ = "STUB: not implemented"
	return
}

func dump(indicator Indicator) (values []float64) { _ = "STUB: not implemented"; return nil }

func indicatorEquals(t *testing.T, expected []float64, indicator Indicator) {
	_ = "STUB: not implemented"
	return
}
