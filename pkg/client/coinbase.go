package client

import (
	"encoding/json"
	"fmt"
	"github.com/SuperMohit/cryto-coins/internal/crypto/entity"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

//https://developers.coinbase.com/
/// coin base related api calls would land here
/// HTTP client to communicate with coinbase
type CoinBaseCaller interface {
	FetchPrices(coinType entity.CoinType) (entity.Crypto, error)
}

//Assignment - create a coinbase account get the token
// Create this client to communicate with the coinbase and fetch btc, eth, polka, doge related data
//

type CoinBaseClient struct {
}

func NewCoinBaseClient() *CoinBaseClient {
	return &CoinBaseClient{}
}

///

/// all the other coins
/// assignment to refactor this so that all the go-routine run in parallel
func (c *CoinBaseClient) FetchPrices(coinType entity.CoinType) (entity.Crypto, error) {
	var wg sync.WaitGroup
	wg.Add(3)

	var spot, sell, buy string

	//spotStr := make(chan string, 1)

	// go routine
	// ANNONYMOUS function running on different goroutine
	// another pattern to run independently using go routine
	go func() {
		spot = Call("https://api.coinbase.com/v2/prices/DOT-USD/spot", "GET", &wg)
	}()

	// read or write in/out channel is blocking in case the channel is unbuffered
	// you need makes some changes

	log.Print("Now printing..........")

	go func() {
		sell = Call("https://api.coinbase.com/v2/prices/DOT-USD/sell", "GET", &wg)
	}()

	go func() {
		buy = Call("https://api.coinbase.com/v2/prices/DOT-USD/buy", "GET", &wg)
	}()

	// 4 + 4 + 4 ~ 12 seconds
	// This is as good as syncronously
	// What should happen   ~ 4 seconds

	// wait-group
	// merge the prices in the below entity
	wg.Wait()

	return entity.NewCrypto(0, 0.0, 0.0, 0.0), nil
}

func Call(url, method string, wg *sync.WaitGroup) string {

	defer wg.Done()
	client := &http.Client{}

	// assignment convert payload to body type
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

		return ""
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	price := entity.Price{}
	err = json.Unmarshal(body, &price)
	if err != nil {
		fmt.Println(err)

		return ""
	}
	//fmt.Println(string(body))

	time.Sleep(4 * time.Second)

	//payload <- price.Data.Amount
	return price.Data.Amount
}
