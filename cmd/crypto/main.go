package main

import "github.com/SuperMohit/cryto-coins/internal/crypto/api"

///  Server  start up code
///  listener code
/// enable versioning in the url path
/// Injection for controller -  service - repository wiring will happen here

/// use embed to read the config file and marsha it into struct
func main() {

	//r := NewRepository
	// inject in the service
	// s := NewCryptoservice(r)
	//  NewCryptoRouter(s)
	api.NewCryptoRouter()

}

/// that microservice will just have some users
/// authentication and authorization
/// JWT for aut and passing tokens amongst services
