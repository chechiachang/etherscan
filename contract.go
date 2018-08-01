package etherscan

import (
	"encoding/json"
	//"fmt"
)

const MODULE_CONTRACT = "contract"

//TODO return struct
func GetContract(contractId string) Contract {
	url := getURL(MODULE_CONTRACT, "getabi", contractId)

	data := Get(url)

	var escaped ContractEscaped
	if err := json.Unmarshal(data, &escaped); err != nil {
		panic(err)
	}

	var result []Result
	if err := json.Unmarshal([]byte(escaped.Result), &result); err != nil {
		panic(err)
	}

	return Contract{
		Status:  escaped.Status,
		Message: escaped.Message,
		Result:  result,
	}
}

type ContractEscaped struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type Contract struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}

type Result struct {
	Constant bool `json:"constant,omitempty"`
	Inputs   []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"inputs,omitempty"`
	Name            string        `json:"name,omitempty"`
	Outputs         []interface{} `json:"outputs,omitempty"`
	Payable         bool          `json:"payable,omitempty"`
	StateMutability string        `json:"stateMutability,omitempty"`
	Type            string        `json:"type"`
	Anonymous       bool          `json:"anonymous,omitempty"`
}
