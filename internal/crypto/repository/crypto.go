package repository

/// interface
/// CRUD
//// to communicate with the database
///https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents

type CryptoRepository interface {
	/// CRUD methods
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
