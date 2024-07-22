package gotech

// SMA calculates the Simple Moving Average of a given period.
func SMA(klines []kline, period int) []Indicator {
	if period > len(klines) {
		return nil
	}
	smas := make([]Indicator, len(klines)-period+1)
	for i := 0; i < len(klines)-period+1; i++ {
		var sum float64
		for j := i; j < i+period; j++ {
			sum += klines[j].Price()
		}
		smas[i] = Indicator{
			Value: sum / float64(period),
			Time:  klines[i+period-1].Time(),
		}
	}
	return smas
}
