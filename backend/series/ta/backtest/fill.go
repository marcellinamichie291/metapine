package backtest

type FillType int

const (
	LIMIT FillType = iota
	MARKET
	STOP
)

/* Fills are Interesting but not mandatory
 */
type Fill struct {
	Side  bool
	Type  FillType
	Price float64
	Size  float64
	Time  int64
}
