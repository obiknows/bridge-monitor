package main

import (
	"fmt"
	"time"

	// "time"

	"github.com/nanmu42/etherscan-api"
)

var ethscan *etherscan.Client
var CTCAddress string
var MaxAllowableTimeDifference time.Duration

func init()  {
	// setup the client
	// 1. create a API client for eth mainnet (hard-coded test API key)
	ethscan = etherscan.New(etherscan.Mainnet, "IBXNW84YSJV17GNXGVHKT6HUS1BXJYXMMC")

	// 2. fill in constants
	CTCAddress = "0x5E4e65926BA27467555EB562121fac00D24E9dD2" // address for the CTC
	MaxAllowableTimeDifference = 2 * time.Minute // use a max 2 minute window

	// startup prompt
	fmt.Println("The Optimism Bridge Monitor 🌉🖥️")
	fmt.Println("--------------------------------")
	fmt.Println("Prints 'True' is bridge is alive.")
	fmt.Println("Prints 'False' if bridge is down.")
	fmt.Print("Prints 'Error' with a message and exits if there is an error with bridge monitoring system.\n\n\n")
}

func main()  {
	// the free etherscan API has a rate-limit of 5 sec/req
	// so... we'll poll every 10 seconds to be quick but safe
	tick := time.Tick(10 * time.Second)
	for range tick {
		alive, err := checkIfOptimismBridgeIsAlive()
		if err != nil {
			fmt.Println("⚠️ Error:", err.Error())
			fmt.Println("Program Exiting.")
			break
		}
		if alive {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}


func checkCTCTransctionsTimeDifference() (time.Duration, error) {
	// ERC20 transactions from/to a specified address
	txs, err := ethscan.NormalTxByAddress(CTCAddress, nil, nil, 1, 5, true)
	if err != nil {
		err = fmt.Errorf("there was an error getting tx's from the CTC. please check your")
		return 0, err
	} 
		// get most recent and oldest tx's from window...
		mostRecentTxTime := txs[0].TimeStamp.Time()
		oldestTxTime := txs[len(txs)-1].TimeStamp.Time()
		// .. calculate the time difference and return
		timeDiff := mostRecentTxTime.Sub(oldestTxTime)

		return timeDiff, nil

}

func checkIfOptimismBridgeIsAlive() (bool, error) {
	timeDiff, err := checkCTCTransctionsTimeDifference()
	if err != nil {
		return false, err
	}

	if (timeDiff > MaxAllowableTimeDifference) {
		return false, err
	}

	return true, err
} 