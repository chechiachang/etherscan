package etherscan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Get(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return contents
}
