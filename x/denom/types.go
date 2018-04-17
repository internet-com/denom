package denom

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// A really cool msg type, these fields are can be entirely arbitrary and
// custom to your message
type ClaimDomainMessage struct {
	Sender     sdk.Address
	DomainName string
	Fee        uint64 // Optional
}

var _ sdk.Msg = ClaimDomainMessage{}

// New cool message
func NewClaimDomainMessage(sender sdk.Address, domainName string, fee uint64) ClaimDomainMessage {
	return ClaimDomainMessage{
		Sender:     sender,
		DomainName: domainName,
		Fee:        fee,
	}
}

// enforce the msg type at compile time

// nolint
func (msg ClaimDomainMessage) Type() string                            { return "denom" }
func (msg ClaimDomainMessage) Get(key interface{}) (value interface{}) { return nil }
func (msg ClaimDomainMessage) GetSigners() []sdk.Address               { return []sdk.Address{msg.Sender} }
func (msg ClaimDomainMessage) String() string {
	return fmt.Sprintf("ClaimDomainMessage{Sender: %v, DomainName: %v}", msg.Sender, msg.DomainName)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg ClaimDomainMessage) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 || len(msg.DomainName) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String()).Trace("")
	}
	return nil
}

// Get the bytes for the message signer to sign on
func (msg ClaimDomainMessage) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

type SetDomainForSaleMessage struct {
	Sender     sdk.Address
	DomainName string
	SalePrice  uint64
}

var _ sdk.Msg = SetDomainForSaleMessage{}

// New cool message
func NewSetDomainForSaleMessage(sender sdk.Address, domainName string, salePrice uint64) SetDomainForSaleMessage {
	return SetDomainForSaleMessage{
		Sender:     sender,
		DomainName: domainName,
		SalePrice:  salePrice,
	}
}

func (msg SetDomainForSaleMessage) Type() string                            { return "denom" }
func (msg SetDomainForSaleMessage) Get(key interface{}) (value interface{}) { return nil }
func (msg SetDomainForSaleMessage) GetSigners() []sdk.Address               { return []sdk.Address{msg.Sender} }
func (msg SetDomainForSaleMessage) String() string {
	return fmt.Sprintf("SetDomainForSaleMessage{Sender: %v, DomainName: %v}", msg.Sender, msg.DomainName)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg SetDomainForSaleMessage) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 || len(msg.DomainName) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String()).Trace("")
	}
	return nil
}

// Get the bytes for the message signer to sign on
func (msg SetDomainForSaleMessage) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// A really cool msg type, these fields are can be entirely arbitrary and
// custom to your message
type ValidateDomainMessage struct {
	Sender     sdk.Address
	DomainName string
	Owner      sdk.Address
	Fee        uint64 // Optional
}

var _ sdk.Msg = ValidateDomainMessage{}

// New cool message
func NewValidateDomainMessage(sender sdk.Address, domainName string, owner sdk.Address) ValidateDomainMessage {
	return ValidateDomainMessage{
		Sender:     sender,
		DomainName: domainName,
		Owner:      owner,
	}
}

// enforce the msg type at compile time

// nolint
func (msg ValidateDomainMessage) Type() string                            { return "denom" }
func (msg ValidateDomainMessage) Get(key interface{}) (value interface{}) { return nil }
func (msg ValidateDomainMessage) GetSigners() []sdk.Address               { return []sdk.Address{msg.Sender} }
func (msg ValidateDomainMessage) String() string {
	return fmt.Sprintf("ClaimDomainMessage{Sender: %v, DomainName: %v}", msg.Sender, msg.DomainName)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg ValidateDomainMessage) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 || len(msg.DomainName) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String()).Trace("")
	}
	return nil
}

// Get the bytes for the message signer to sign on
func (msg ValidateDomainMessage) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
