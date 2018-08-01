package etherscan

const MODULE_CONTRACT = "contract"

//TODO return struct
func GetContract(contractId string) string {
	url := getURL(MODULE_CONTRACT, "getabi", contractId)
	return Get(url)
}
