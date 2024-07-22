package gotech

type kline interface {
	Price() float64
	Time() int64
}
