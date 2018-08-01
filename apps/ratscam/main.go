package main

import (
	"fmt"
	"github.com/chechiachang/etherscan"
)

const RATSCAM_CONTRACT = "0x8a883a20940870Dc055F2070ac8eC847ed2d9918"

func main() {
	fmt.Println("ratscam...")
	fmt.Println(etherscan.GetContract(RATSCAM_CONTRACT))
}
