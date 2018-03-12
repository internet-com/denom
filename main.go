package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/svaishnavy/denom/denom"
	"github.com/tendermint/abci/server"
	"github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"
)

/*
   RegisterDomain(domainName, transferTime, askingPrice, encryptedEmailAddress)
   Verified(domainName, owner) // Done by validators
   BidDomain(domainName, transferBy, bidPrice)
   ProposeAlterBid(domainName, transferBy, bidPrice)
   AcceptAlterBid(domainName)
   AcceptBid(domainName)
   Deposit(domainName, value)
   Withdraw(domainName) // Withdraw after transfered  date + 3 days.
   Transferd(domainName) // Seller and Buyer updates with transferd message, and money is sent to seller.
   Dispute(domainName) // Dispute
   AddArbitrer(domainName)
   AcceptArbitrer(domainName)
   ReleaseFunds(domainName, buyer, seller) // Done by Arbiter
*/

func main() {
	addrPtr := flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	abciPtr := flag.String("abci", "socket", "socket | grpc")
	//persistencePtr := flag.String("persist", "", "directory to use for a database")
	flag.Parse()

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Create the application - in memory or persisted to disk
	var app types.Application
	app = denom.NewDenomApp(true)

	// Start the listener
	srv, err := server.NewServer(*addrPtr, *abciPtr, app)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	srv.SetLogger(logger.With("module", "abci-server"))
	if err := srv.Start(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})
	domain := flag.String("domain", "denom.org", "Enter the domain name")
	flag.Parse()
	s, _ := net.LookupTXT(*domain)
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}
