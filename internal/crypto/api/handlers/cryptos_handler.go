package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/SuperMohit/cryto-coins/internal/crypto/usecase"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type CryptoHandler struct {
	service      usecase.CryptoFetcher
	orderService usecase.CryptoService
}

func NewCryptoHandler(service usecase.CryptoFetcher) *CryptoHandler {
	return &CryptoHandler{service: service}
}

/// variable
/// Refactor  the common code
/// hander -   just contains logic to read request and write to the response.
//  move third party calls to coinbase.go

///  transfer to the service code

func (c *CryptoHandler) CryptosPolkaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cryptos POLKA service invoked.")

	// Move below code to the client
	//url := "https://api.coinbase.com/v2/prices/POLKA-SGD/spot"
	//method := "GET"
	//
	//client := &http.Client {
	//}
	//req, err := http.NewRequest(method, url, nil)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	var response map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func Cryptosdogehandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cryptos DOGE service invoked.")

	url := "https://api.coinbase.com/v2/prices/DOGE-USD/spot"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "__cf_bm=3kfPoI39JOSyZ5Sg6ckUhUwVii8i2Shshz5GSf7biiw-1633445063-0-AYHL/NIhbWY3OlgvbVvhheRyubV/YqJz3ISoXpRgb8/1Qdykyx29RLSyvi/cbYYPUveqGR4tJe10WyhWqiNDkMg=; coinbase_device_id=859c49f8-a78f-4aca-b261-3c855c505ed8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var response map[string]interface{}
	json.Unmarshal([]byte(string(body)), &response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func Cryptosethandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cryptos ETH Service invoked.")

	url := "https://api.coinbase.com/v2/prices/ETH-USD/spot"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "__cf_bm=3kfPoI39JOSyZ5Sg6ckUhUwVii8i2Shshz5GSf7biiw-1633445063-0-AYHL/NIhbWY3OlgvbVvhheRyubV/YqJz3ISoXpRgb8/1Qdykyx29RLSyvi/cbYYPUveqGR4tJe10WyhWqiNDkMg=; coinbase_device_id=859c49f8-a78f-4aca-b261-3c855c505ed8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(string(body))

	var response map[string]interface{}
	json.Unmarshal([]byte(string(body)), &response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /cryptos/btc getBTCPrices
// Gets BTC prices of sell, buy and spot
//
// security:
// - apiKey: []
// responses:
//  200: Crypto
func (c *CryptoHandler) CryptosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cryptos Service invoked.")
	urlParams := r.URL.Query()
	log.Println("Header Key: ", urlParams.Get("key"))
	params := mux.Vars(r)
	id, found := mux.Vars(r)["key"]
	log.Println("Request Parameters: ", params)
	log.Println("ID: ", id, "  ", found)

	// write the btc response to the response stream
	btc, err := c.service.GetBTCPrices()

	if err != nil {
		// update the response status to 400
		// write error message
		w.WriteHeader(http.StatusBadRequest)
		// return
	}

	//var response map[string]interface{}
	//json.Unmarshal([]byte(string(body)), &response)
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.WriteHeader(200)

	json.NewEncoder(w).Encode(btc)

	if err != nil {
		log.Println("Error: ", err)
	}
	return
}

// Fill the handler
func (c *CryptoHandler) PlaceBuyOrder(w http.ResponseWriter, r *http.Request) {
	// Get the data from the request body and pass in below function

	// url you will read coinType {coin}
	//c.orderService.BuyCrypto()
}
