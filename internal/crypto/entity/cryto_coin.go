package entity

///
//// struct for coin related data

/// enum for coin types

type CoinType int




type Crypto struct {
	coinType CoinType
	buyPrice float64
	sellPrice float64
	spotPrice float64
}

func NewCrypto(coinType CoinType, buyPrice, sellPrice, spotPrice float64) Crypto {
	return Crypto{
		coinType:  coinType,
		buyPrice:  buyPrice,
		sellPrice: sellPrice,
		spotPrice: spotPrice,
	}
}
