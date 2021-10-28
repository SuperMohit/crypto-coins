package entity

///
//// struct for coin related data

/// enum for coin types
//// enum for 4 crypto types BTC ETH POLKA DOGE
type CoinType int

const (
	BTC CoinType = iota + 1
	ETH
	POLKA
	DOGE
)

// swagger:model Crypto
type Crypto struct {
	// Buy price of the coin
	BuyPrice float64
	// Buy price of the coin
	SellPrice float64
	// Sell price of the coin
	SpotPrice float64
}

func NewCrypto(coinType CoinType, buyPrice, sellPrice, spotPrice float64) Crypto {
	return Crypto{
		BuyPrice:  buyPrice,
		SellPrice: sellPrice,
		SpotPrice: spotPrice,
	}
}

type Price struct {
	Data Data
}

type Data struct {
	Base     string
	Currency string
	Amount   string
}
