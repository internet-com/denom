package registration

type Domain struct {
	ClaimedBy      map[string]uint64
	ValidatedBy    map[string]string
	Owner          string
	ValidatedBlock uint64
	RewardClaimed  uint64
}

type Keeper struct {
}

func (k Keeper) RegisterDomain() {

}

func (k Keeper) ValidateDomain() {

}

func (k Keeper) WithdrawReward() {

}
