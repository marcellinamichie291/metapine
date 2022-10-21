package backtest

/*
SafeFloat
Pretty often indicator have no Value while others have it. To still store everything in a File we have this struct.
SafeFloat is used when we use filters for our Strategies
*/
type SafeFloat struct {
	Safe  bool
	Value float64
}
