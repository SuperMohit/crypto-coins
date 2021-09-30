package usecase

import (
	"github.com/SuperMohit/cryto-coins/internal/crypto/repository"
	"github.com/SuperMohit/cryto-coins/pkg/client"
)

///
//// bussiness code
//// interface    /// -- struct

/// CRUD to communicate with DB layer

type CryptoFetcher interface {
}

type CryptoService struct {
	repo           repository.CryptoRepository
	coinbaseClient client.CoinBaseClient
}

////
////  1st lets only talk to coinbase API
/// client first coin base

func NewCryptoFetcher(repo repository.CryptoRepository) *CryptoFetcher {
	// populate CryptoService
	//  CryptoService{
	//  repo : repo
	//}

	return nil

}
