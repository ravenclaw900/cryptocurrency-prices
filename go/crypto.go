package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cryptoResp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,dogecoin&vs_currencies=usd&include_24hr_change=true")
	if err != nil {
		log.Fatal(err)
	}

	defer cryptoResp.Body.Close()

	type CData struct {
		Price  float64 `json:"usd"`
		Change float64 `json:"usd_24h_change"`
	}

	type Crypto struct {
		Bitcoin  CData `json:"bitcoin"`
		Ethereum CData `json:"ethereum"`
		Dogecoin CData `json:"dogecoin"`
	}

	cryptoData := Crypto{}
	cryptoBytes, err := ioutil.ReadAll(cryptoResp.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(cryptoBytes, &cryptoData)

	fmt.Println("Bitcoin is currently worth $" + strconv.FormatFloat(cryptoData.Bitcoin.Price, 'f', -1, 64) + ", and it has changed by " + strconv.FormatFloat(cryptoData.Bitcoin.Change, 'f', -1, 64) + "% over the past 24 hours.")
	fmt.Println("Ethereum is currently worth $" + strconv.FormatFloat(cryptoData.Ethereum.Price, 'f', -1, 64) + ", and it has changed by " + strconv.FormatFloat(cryptoData.Ethereum.Change, 'f', -1, 64) + "% over the past 24 hours.")
	fmt.Println("Dogecoin is currently worth $" + strconv.FormatFloat(cryptoData.Dogecoin.Price, 'f', -1, 64) + ", and it has changed by " + strconv.FormatFloat(cryptoData.Dogecoin.Change, 'f', -1, 64) + "% over the past 24 hours.")
}
