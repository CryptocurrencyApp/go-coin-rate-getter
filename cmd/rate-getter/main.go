package main

import (
	"os"
	"fmt"
	"encoding/json"
	"bytes"

	getter "github.com/CryptocurrencyApp/go-coin-rate-getter"
)

const newestResultFilePath = "./ratelog/newest/newest.json"

func main() {
	response := getter.Access()

	newestFile, err := os.OpenFile(newestResultFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	newestResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	json.Indent(buffer, newestResponse, "", "    ")

	fmt.Fprintln(newestFile, buffer.String())
}
