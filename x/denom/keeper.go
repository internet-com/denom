package denom

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Cool genesis state, containing the genesis trend
type GenesisState struct {
	trend string
}

// Keeper - handlers sets/gets of custom variables for your module
type Keeper struct {
	ck                 bank.CoinKeeper
	domainStoreKey     sdk.StoreKey
	stakeStore         sdk.StoreKey
	validationStoreKey sdk.StoreKey
	userStoreKey       sdk.StoreKey

	storeKey sdk.StoreKey // The (unexposed) key used to access the store from the Context.
}

type Domain struct {
	Owner          sdk.Address
	ValidatedBy    map[string]string
	ValidatedBlock uint64
	ClaimedBy      map[string]bool
}

func (k Keeper) GetBondedTokens(ctx sdk.Context, address sdk.Address) uint64 {
	//store := ctx.KVStore(k.stakeStore)
	//addressBytes := []byte(address.String())
	//bondBytes := store.Get(addressBytes)
	return 1
}

// NewKeeper - Returns the Keeper
func NewKeeper(key sdk.StoreKey, bankKeeper bank.CoinKeeper) Keeper {
	return Keeper{ck: bankKeeper, storeKey: key}
}

// GetDomain - Returns the domain details from store
func (k Keeper) GetDomain(ctx sdk.Context, domainName string) (Domain, error) {
	store := ctx.KVStore(k.storeKey)
	domainNameBytes := []byte(domainName)
	if store.Has(domainNameBytes) {
		domainBytes := store.Get(domainNameBytes)
		buf := new(bytes.Buffer)
		buf.Read(domainBytes)
		domain := Domain{}
		err := binary.Read(buf, binary.BigEndian, &domain)
		return domain, err
	} else {
		return Domain{ClaimedBy: map[string]bool{}, ValidatedBy: map[string]string{}}, nil
	}
}

// Implements sdk.AccountMapper.
func (k Keeper) Claim(ctx sdk.Context, sender sdk.Address, domainName string) {
	domain, err := k.GetDomain(ctx, domainName)
	if err == nil {
		store := ctx.KVStore(k.storeKey)
		domainNameBytes := []byte(domainName)
		domain.ClaimedBy[sender.String()] = true
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, &domain)
		store.Set(domainNameBytes, buf.Bytes())
	}
}

// Implements sdk.AccountMapper.
func (k Keeper) Validate(ctx sdk.Context, sender sdk.Address, domainName string, owner sdk.Address) {
	domain, err := k.GetDomain(ctx, domainName)
	if err == nil {
		store := ctx.KVStore(k.storeKey)
		domainNameBytes := []byte(domainName)
		// Check if sender is validator and total votes for the claimer is > 2/3rd total bonded DNOM
		domain.ValidatedBy[sender.String()] = owner.String()
		votes := uint64(0)
		for validator, voted_for := range domain.ValidatedBy {
			if owner.String() == voted_for {
				validatorAddress, err := sdk.GetAddress(validator)
				if err == nil {
					votes += k.GetBondedTokens(ctx, validatorAddress)
				}
			}
		}
		// Modify the below line to check the bonded stake.
		if votes >= 1 {
			domain.ClaimedBy = map[string]bool{}
			domain.ValidatedBy = map[string]string{}
			domain.Owner = owner
			k.ck.AddCoins(ctx, owner, []sdk.Coin{sdk.Coin{Amount: 1000000000, Denom: "DNOM"}})
		}
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, &domain)
		store.Set(domainNameBytes, buf.Bytes())
	}
}

// InitGenesis - store the genesis trend
func (k Keeper) InitGenesis(ctx sdk.Context, data json.RawMessage) error {
	var state GenesisState
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}
	//k.setTrend(ctx, state.trend)
	return nil
}
