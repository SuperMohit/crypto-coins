package api

import (
	"context"
	"github.com/SuperMohit/cryto-coins/internal/crypto/api/handlers"
	"github.com/SuperMohit/cryto-coins/internal/crypto/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/// all the routes specific to crypto
//// and handlers as well

//// GET    /cryptos     --- all cryptos
///  GET    /cryptos/btc  ---- BTC related data
///  GET    /cryptos/eth  ---- ETH related data
///  GET    /cryptos/doge  ---- DOGE related data
///  GET    /cryptos/polka  ---- POLKA related data

const (
	maxHeaderBytes    = 1024
	readHeaderTimeout = 0o3
	readTimeout       = 0o3
	writeTimeout      = 15
	idleTimeout       = 60
	graceTime         = 15
)

type Crypto struct {
	serv usecase.CryptoFetcher
}

type CryptoRouter interface {
}

func NewCryptoRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/cryptos/btc", handlers.Cryptoshandler).Methods("GET")
	router.HandleFunc("/cryptos/eth", handlers.Cryptosethandler).Methods("GET")
	router.HandleFunc("/cryptos/doge", handlers.Cryptosdogehandler).Methods("GET")
	router.HandleFunc("/cryptos/polka", handlers.Cryptospolkahandler).Methods("GET")
	listen(":9001", router)
}

func listen(address string, handler http.Handler) {
	server :=&http.Server{
		Addr:              address,
		Handler:           handler,
		ReadTimeout:       readTimeout * time.Second,
		ReadHeaderTimeout: readHeaderTimeout * time.Second,
		WriteTimeout:      writeTimeout * time.Second,
		IdleTimeout:       idleTimeout * time.Second,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Println("Started and Listening at address: ", address)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("Error and Shutting down")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<- c

	log.Println("Shutting down")

	ctx,cancel := context.WithTimeout(context.Background(), graceTime * time.Second)
	defer cancel()

	if err :=server.Shutdown(ctx);err!= nil {
		log.Println("Error resulted in shutdown.")

	}
}
