package gotech

const (
	ClosePrice = iota
	OpenPrice
	HighPrice
	LowPrice
	TypicalPrice
	WeightedPrice
)

type Kline struct {
	Open  float64
	High  float64
	Low   float64
	Close float64

	OpenTime  int64
	CloseTime int64
}

func (s Kline) Typical() float64 {
	return (s.High + s.Low + s.Close) / 3
}

func (s Kline) Weighted() float64 {
	return (s.High + s.Low + s.Close*2) / 4
}

func (s Kline) Price(priceType ...int) float64 {
	if len(priceType) == 0 {
		return s.Close
	}

	switch priceType[0] {
	case ClosePrice:
		return s.Close
	case OpenPrice:
		return s.Open
	case HighPrice:
		return s.High
	case LowPrice:
		return s.Low
	case TypicalPrice:
		return s.Typical()
	case WeightedPrice:
		return s.Weighted()
	default:
		return s.Close
	}
}

func (s Kline) Time() int64 {
	return s.CloseTime
}
