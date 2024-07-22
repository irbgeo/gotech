package gotech

import "math"

// RSI calculates the Relative Strength Index of a given period.
func RSI(klines []Kline, period int) []Indicator {
	if period >= len(klines) {
		return nil
	}

	rsis := make([]Indicator, len(klines)-period)
	var gain, loss float64

	// Calculate initial gain and loss
	for i := 1; i <= period; i++ {
		diff := klines[i].Price() - klines[i-1].Price()
		if diff > 0 {
			gain += diff
		} else {
			loss -= diff
		}
	}

	// Calculate first RSI
	avgGain := gain / float64(period)
	avgLoss := loss / float64(period)
	var rsi float64
	if avgLoss == 0 {
		rsi = 100
	} else {
		rs := avgGain / avgLoss
		rsi = 100 - (100 / (1 + rs))
	}
	rsis[0] = Indicator{
		Value: rsi,
		Time:  klines[period].CloseTime,
	}

	// Calculate subsequent RSIs
	for i := period + 1; i < len(klines); i++ {
		diff := klines[i].Price() - klines[i-1].Price()
		gain = (avgGain*float64(period-1) + math.Max(diff, 0)) / float64(period)
		loss = (avgLoss*float64(period-1) + math.Max(-diff, 0)) / float64(period)

		if loss == 0 {
			rsi = 100
		} else {
			rs := gain / loss
			rsi = 100 - (100 / (1 + rs))
		}

		rsis[i-period] = Indicator{
			Value: rsi,
			Time:  klines[i].CloseTime,
		}

		avgGain = gain
		avgLoss = loss
	}

	return rsis
}
