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
	ck bank.CoinKeeper

	storeKey sdk.StoreKey // The (unexposed) key used to access the store from the Context.
}

// NewKeeper - Returns the Keeper
func NewKeeper(key sdk.StoreKey, bankKeeper bank.CoinKeeper) Keeper {
	return Keeper{bankKeeper, key}
}

type Domain struct {
	ValidatedBy map[string]bool
	Owner       sdk.Address
	ClaimedBy   map[string]map[string]bool
}

// GetTrend - returns the current cool trend
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
		return Domain{ClaimedBy: map[string]map[string]bool{}}, nil
	}
}

// Implements sdk.AccountMapper.
func (k Keeper) Claim(ctx sdk.Context, sender sdk.Address, domainName string) {
	domain, err := k.GetDomain(ctx, domainName)
	if err == nil {
		store := ctx.KVStore(k.storeKey)
		domainNameBytes := []byte(domainName)
		domain.ClaimedBy[sender.String()] = map[string]bool{}
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, &domain)
		store.Set(domainNameBytes, buf.Bytes())
	}
}

// Implements sdk.AccountMapper.
func (k Keeper) Validate(ctx sdk.Context, sender sdk.Address, domainName string) {
	domain, err := k.GetDomain(ctx, domainName)
	if err == nil {
		store := ctx.KVStore(k.storeKey)
		domainNameBytes := []byte(domainName)
		domain.ClaimedBy[sender.String()] = map[string]bool{}
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
