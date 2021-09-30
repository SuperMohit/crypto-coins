package api

import "github.com/SuperMohit/cryto-coins/internal/crypto/usecase"

/// all the routes specific to crypto
//// and handlers as well

//// GET    /cryptos     --- all cryptos
///  GET    /cryptos/btc  ---- BTC related data
///  GET    /cryptos/eth  ---- ETH related data
///  GET    /cryptos/doge  ---- DOGE related data
///  GET    /cryptos/polka  ---- POLKA related data

type Crypto struct {
	serv usecase.CryptoFetcher
}

type CryptoRouter interface {
}

func NewCryptoRouter() {

}
