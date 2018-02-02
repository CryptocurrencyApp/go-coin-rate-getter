package rate_getter

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"fmt"
)

type response struct {
}

const baseUrl = "https://api.coinmarketcap.com/v1/ticker"
const coinTypeCount = 300

func Access() {
	result := request()
	fmt.Println(result)
}

func request() string {
	option := createOption()

	resp, err := http.Get(baseUrl + "?" + option.Encode())
	if err != nil {
		panic(err)
	}

	result, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	return string(result)
}

func createOption() url.Values {
	value := url.Values{}
	value.Add("start", "1")
	value.Add("limit", string(coinTypeCount))
	value.Add("convert", "JPY")

	return value
}
