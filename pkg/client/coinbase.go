package client

import (
	"fmt"
	"github.com/SuperMohit/cryto-coins/internal/crypto/entity"
	"io/ioutil"
	"net/http"
)

//https://developers.coinbase.com/
/// coin base related api calls would land here
/// HTTP client to communicate with coinbase
type CoinBaseClient interface {
}

//Assignemnt - create a coinbase account get the token
// Create this client to communicate with the coinbase and fetch btc, eth, polka, doge related data
//



/// all the other coins
func FetchPrices(coinType entity.CoinType) entity.Crypto {
	var spot, sell, buy float64

	//go Call()  --  spot

	//go Call()  -- sell

	//go Call  -- buy


	// waitgroup

	return entity.NewCrypto(0, buy, sell, spot)
}








func Call(url, method string, payload interface{}) {

	client := &http.Client {
	}

	// assignment convert payload to body type
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

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
}
