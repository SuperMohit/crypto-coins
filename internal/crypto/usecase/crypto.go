package usecase

import (
	"github.com/SuperMohit/cryto-coins/internal/crypto/entity"
	"github.com/SuperMohit/cryto-coins/internal/crypto/repository"
	"github.com/SuperMohit/cryto-coins/pkg/client"
	"go.uber.org/zap"
)

///
//// bussiness code
//// interface    /// -- struct

/// CRUD to communicate with DB layer

type CryptoFetcher interface {
	GetBTCPrices() (entity.Crypto, error)
}

type CryptoBuyer interface {
	BuyCrypto(coinType entity.CoinType, buyPrice float64) error
}

type CryptoService struct {
	repo           repository.CryptoRepository
	coinbaseClient client.CoinBaseCaller
	log            *zap.SugaredLogger
}

////
////  1st lets only talk to coinbase API
/// client first coin base
///  inject the client
func NewCryptoFetcher(repo repository.CryptoRepository, coinbaseClient client.CoinBaseCaller,
	log *zap.SugaredLogger) *CryptoService {
	// populate CryptoService
	//  CryptoService{
	//  repo : repo
	//}

	return &CryptoService{
		repo:           repo,
		coinbaseClient: coinbaseClient,
		log:            log,
	}

}

func (c *CryptoService) GetBTCPrices() (entity.Crypto, error) {
	btc, err := c.coinbaseClient.FetchPrices(entity.BTC)
	if err != nil {
		c.log.Errorf("recieved error %v", err)
	}
	return btc, err
}

func (c *CryptoService) BuyCrypto(coinType entity.CoinType, buyPrice float64) error {
	c.log.Info("Started buying cryptos")

	// pass the params as the values
	return c.repo.SaveOrder(entity.BuyOrder{
		CoinType: 0,
		BuyPrice: 0,
		LotSize:  0,
	})

}
