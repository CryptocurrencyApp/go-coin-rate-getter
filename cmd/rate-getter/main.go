package main

import (
	"fmt"

	getter "github.com/CryptocurrencyApp/go-coin-rate-getter"
)

func main() {
	response := getter.Access()
	fmt.Println(response)
}
