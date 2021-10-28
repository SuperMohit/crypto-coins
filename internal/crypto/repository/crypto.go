package repository

import "github.com/SuperMohit/cryto-coins/internal/crypto/entity"

/// interface
/// CRUD
//// to communicate with the database

type CryptoRepository interface {
	/// CRUD methods
	SaveOrder(buyOrder entity.BuyOrder) error
}

// this will implement all the methods of CryptoRepository
type CryptoStorage struct {
	// db - client
}

func NewCryptoRepository() *CryptoRepository {
	///
	/// initialize crypto storage
	return nil
}

/// p
func (cs *CryptoStorage) SaveOrder(buyOrder entity.BuyOrder) error {
	/// here you will make a insert call to the mongo db
	///https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents

	return nil
}
