package main

import (
	"context"
	"github.com/SuperMohit/cryto-coins/internal/crypto/api"
	"github.com/SuperMohit/cryto-coins/internal/crypto/api/handlers"
	"github.com/SuperMohit/cryto-coins/internal/crypto/repository"
	"github.com/SuperMohit/cryto-coins/internal/crypto/usecase"
	"github.com/SuperMohit/cryto-coins/pkg/client"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"time"
)

const (
	maxHeaderBytes    = 1024
	readHeaderTimeout = 0o3
	readTimeout       = 0o3
	writeTimeout      = 15
	idleTimeout       = 60
	graceTime         = 15
)

///configmap
///secret
///sealed-secret

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Info("started the server")

	repository.NewCryptoRepository()
	client := client.NewCoinBaseClient()

	// create a new coinbase client
	service := usecase.NewCryptoFetcher(nil, client, sugar)

	handlers := handlers.NewCryptoHandler(service)

	router := api.NewCrypto(handlers)

	r := router.CryptoRouter()

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)

	sugar.Info(reflect.TypeOf(router))

	r.Handle("/docs", sh)
	listen(":9001", r)

}

/// that microservice will just have some users
/// authentication and authorization
/// JWT for aut and passing tokens amongst services

func listen(address string, handler http.Handler) {
	server := &http.Server{
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
	<-c

	log.Println("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), graceTime*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error resulted in shutdown.")

	}
}
