package main

import (
	"fmt"

	"github.com/nanmu42/etherscan-api"
)

func main()  {
	fmt.Println("hello world, its me the bridge monitor")
	fmt.Println("checking an account balance as a test....")
	// create a API client for specified ethereum net
	// there are many pre-defined network in package
	ethscan := etherscan.New(etherscan.Mainnet, "IBXNW84YSJV17GNXGVHKT6HUS1BXJYXMMC")

	// check account balance
	balance, err := ethscan.AccountBalance("0x281055afc982d96fab65b3a49cac8b878184cb16")
	if err != nil {
		panic(err)
	}
	// balance in wei, in *big.Int type
	fmt.Println("account balance:", balance.Int())

	
}


//Test API Key: IBXNW84YSJV17GNXGVHKT6HUS1BXJYXMMC