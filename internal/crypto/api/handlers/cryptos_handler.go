package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func Cryptospolkahandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cryptos POLKA service invoked.")

	url := "https://api.coinbase.com/v2/prices/POLKA-SGD/spot"
	method := "GET"

	client := &http.Client {
	}
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

	var response map[string]interface{}
	json.Unmarshal([]byte(string(body)), &response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func Cryptosdogehandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cryptos DOGE service invoked.")

	url := "https://api.coinbase.com/v2/prices/DOGE-USD/spot"
	method := "GET"

	client := &http.Client {
	}
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

	client := &http.Client {
	}
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

func Cryptoshandler(w http.ResponseWriter, r *http.Request)  {
	log.Println("Cryptos Service invoked.")
	urlParams := r.URL.Query()
	log.Println("Header Key: ", urlParams.Get("key"))
	params := mux.Vars(r)
	id, found := mux.Vars(r)["key"]
	log.Println("Request Parameters: ", params)
	log.Println("ID: ", id, "  ", found)

	url := "https://api.coinbase.com/v2/prices/BTC-USD/spot"
	method := "GET"

	client := &http.Client {
	}
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

	var response map[string]interface{}
	json.Unmarshal([]byte(string(body)), &response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}