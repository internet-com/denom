package arbitrer

type Arbitrer struct {
	Name              string
	Website           string
	MinArbitrationFee uint64
	MaxArbitrationFee uint64
	ArbitrationFee    uint16
	BuyerVotes        uint16
	SellerVotes       uint16
	TotalTransactions uint16
}

type Keeper struct {
}

func (k Keeper) AddRating() {

}

func (k Keeper) SetArbitrationFee() {

}

func (k Keeper) Register() {

}
