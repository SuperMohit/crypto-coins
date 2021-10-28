package api

import (
	"github.com/SuperMohit/cryto-coins/internal/crypto/api/handlers"
	"github.com/gorilla/mux"
)

/// all the routes specific to crypto
//// and handlers as well

//// GET    /cryptos     --- all cryptos
///  GET    /cryptos/btc  ---- BTC related data
///  GET    /cryptos/eth  ---- ETH related data
///  GET    /cryptos/doge  ---- DOGE related data
///  GET    /cryptos/polka  ---- POLKA related data

type Crypto struct {
	handler *handlers.CryptoHandler
}

func NewCrypto(handler *handlers.CryptoHandler) *Crypto {
	return &Crypto{handler: handler}
}

type CryptoRouter interface {
}

//  Crypto Api:
//   version: 0.0.1
//   title: Crypto API for buying and selling Crypto coins
//  Schemes: http
//  Host: localhost:9001
//  BasePath: /
//  Produces:
//    - application/json
//
// securityDefinitions:
//  apiKey:
//    type: apiKey
//    in: header
//    name: authorization
// swagger:meta

func (c *Crypto) CryptoRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/cryptos/btc", c.handler.CryptosHandler).Methods("GET")

	// change these as well as member handler functions
	router.HandleFunc("/cryptos/eth", handlers.Cryptosethandler).Methods("GET")

	//// change these as well as member handler functions
	router.HandleFunc("/cryptos/doge", handlers.Cryptosdogehandler).Methods("GET")
	router.HandleFunc("/cryptos/polka", c.handler.CryptosPolkaHandler).Methods("GET")

	/// assignment -2 to finish the save to DB flow.
	router.HandleFunc("/cryptos/{coin}", c.handler.PlaceBuyOrder)

	return router
}
