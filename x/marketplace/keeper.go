package marketplace

type DomainItem struct {
	MinimumPrice   uint64
	TimeToTransfer uint8
	Bids           map[string]DomainBid
}

type DomainBid struct {
	SealedBid string
	Price     uint64
	TimeBy    uint64
	Arbitrer  string
}

type Keeper struct {
}

func (k Keeper) RegisterDomain() {

}

func (k Keeper) ValidateDomain() {
}

func (k Keeper) SetDomainForSale() {
}
