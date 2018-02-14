package rate_getter

import (
	"net/url"
	"net/http"
	"io"
	"io/ioutil"
	"os"
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
		InfoList: nil,
	}

	if err := json.Unmarshal(rawResult, &response.InfoList); err != nil {
		panic(err)
	}

	return response
}

func Archive() error {
	// Open source file
	srcPath := "/home/ubuntu/ratelog/newest/newest.json"

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return err
	}

	// newest/newest.jsonの中身が空の場合（初めてのアクセスの時）はスルーする
	if len(bytes) == 0 {
		return nil
	}

	result := Response{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}
	defer srcFile.Close()

	// Open destination file
	dstPath := "/home/ubuntu/ratelog/archive/" + result.GetAt.Format(time.RFC3339) + ".json"
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
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
