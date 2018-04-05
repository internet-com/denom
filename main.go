package main

import (
	"flag"
	"fmt"
	"net"
)

/*
   RegisterDomain(domainName, transferTime, askingPrice, encryptedEmailAddress)
   Verified(domainName, owner) // Done by validators
   BidDomain(domainName, transferBy, bidPrice)
   ListBids(domainName) -> Bid(domain, price, transferBefore)
   ProposeAlterBid(domainName, transferBy, bidPrice)
   AcceptAlterBid(domainName)
   AcceptBid(domainName)
   Deposit(domainName, value)
   Withdraw(domainName) // Withdraw after transfered  date + 3 days.
   Transfered(domainName) // Seller and Buyer updates with transfered message, and money is sent to seller.
   Dispute(domainName) // Dispute
   AddArbitrer(domainName)
   AcceptArbitrer(domainName)
   ReleaseFunds(domainName, buyer, seller) // Done by Arbiter

   RequestDomainForSale(domainName, optional:emailName)

   {from: 0x, to: 0x, data: , value: fee, nonce: }
*/

func main() {
	domain := flag.String("domain", "denom.org", "Enter the domain name")
	flag.Parse()
	s, _ := net.LookupTXT(*domain)
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}
