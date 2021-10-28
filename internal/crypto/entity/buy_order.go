package entity

// anything else you can add more fields
type BuyOrder struct {
	CoinType CoinType
	BuyPrice float64
	LotSize  int
}
