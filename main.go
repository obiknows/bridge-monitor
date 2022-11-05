package main

import (
	"fmt"
	"time"

	"github.com/nanmu42/etherscan-api"
)

var ethscan *etherscan.Client

func init()  {
	// setup the client
	// 1. create a API client for eth mainnet (hard-coded test API key)
	ethscan = etherscan.New(etherscan.Mainnet, "IBXNW84YSJV17GNXGVHKT6HUS1BXJYXMMC")

	// startup prompt
	fmt.Println("hello world, its me the bridge monitor")
	fmt.Println("checking an account balance as a test....")
}

func main()  {
	// the free etherscan API has a rate-limit of 5 sec/req
	// so... we'll poll every 7 seconds to be quick but safe
	tick := time.Tick(7 * time.Second)
	for range tick {
			fmt.Println("7 second Tick")
			checkAccountBalance()
			checkIfOptimismBridgeIsAlive()
	}
}

func checkAccountBalance() {
	// check account balance
	balance, err := ethscan.AccountBalance("0x281055afc982d96fab65b3a49cac8b878184cb16")
	if err != nil {
		panic(err)
	}
	// balance in wei, in *big.Int type
	fmt.Println("account balance:", balance.Int())

}

func checkIfOptimismBridgeIsAlive() bool {
	return true
} 