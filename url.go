package etherscan

import (
	"strings"
)

// https://api.etherscan.io/api?module=contract\&action=getabi\&address=0x8a883a20940870Dc055F2070ac8eC847ed2d9918\&apikey=${API_TOKEN}
func getURL(module, action, address string) string {
	url := []string{}
	url = append(url, ETHERSCAN_URL)
	url = append(url, "?")
	url = append(url, "module=")
	url = append(url, module)
	url = append(url, "&")
	url = append(url, "action=")
	url = append(url, action)
	url = append(url, "&")
	url = append(url, "address=")
	url = append(url, address)
	url = append(url, "&")
	url = append(url, "apikey=")
	url = append(url, getApiKey())

	return strings.Join(url, "")
}
