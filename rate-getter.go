package rate_getter

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
	"strconv"
)

type Response struct {
	GetAt    time.Time
	InfoList []CoinInfo
}

type CoinInfo struct {
	Id               string      `json:"id"`
	Name             string      `json:"name"`
	Symbol           string      `json:"symbol"`
	PriceUsd         json.Number `json:"price_usd"`
	PriceJpy         json.Number `json:"price_jpy"`
	PriceBtc         json.Number `json:"price_btc"`
	PercentChange1h  json.Number `json:"percent_change_1h"`
	PercentChange24h json.Number `json:"percent_change_24h"`
	PercentChange7d  json.Number `json:"percent_change_7d"`
}

const baseUrl = "https://api.coinmarketcap.com/v1/ticker/"
const coinTypeCount = 300

func Access() Response {
	rawResult := request()

	response := Response{
		GetAt:    time.Now(),
		InfoList: []CoinInfo{},
	}

	if err := json.Unmarshal(rawResult, &response.InfoList); err != nil {
		panic(err)
	}

	return response
}

func request() []byte {
	option := createOption()

	resp, err := http.Get(baseUrl + "?" + option.Encode())
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	return result
}

func createOption() url.Values {
	value := url.Values{}
	value.Add("start", "0")
	value.Add("limit", strconv.Itoa(coinTypeCount))
	value.Add("convert", "JPY")

	return value
}
