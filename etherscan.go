package etherscan

import (
	"os"
)

const ETHERSCAN_URL = "https://api.etherscan.io/api"

var (
	apiKey string
)

func init() {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		panic("no api key")
	}
}

func getApiKey() string {
	return apiKey
}
