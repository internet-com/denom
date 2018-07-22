package registration

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type RegisterDomainMessage struct {
	Sender     sdk.AccAddress
	DomainName string
	Fee        uint64 // Optional
}

var _ sdk.Msg = RegisterDomainMessage{}

func NewRegisterDomainMessage(sender sdk.AccAddress, domainName string, fee uint64) RegisterDomainMessage {
	return RegisterDomainMessage{
		Sender:     sender,
		DomainName: domainName,
		Fee:        fee,
	}
}

// enforce the msg type at compile time

// nolint
func (msg RegisterDomainMessage) Type() string                            { return "register" }
func (msg RegisterDomainMessage) Get(key interface{}) (value interface{}) { return nil }
func (msg RegisterDomainMessage) GetSigners() []sdk.AccAddress            { return []sdk.AccAddress{msg.Sender} }
func (msg RegisterDomainMessage) String() string {
	return fmt.Sprintf("RegisterDomainMessage{Sender: %v, DomainName: %v}", msg.Sender, msg.DomainName)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg RegisterDomainMessage) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 || len(msg.DomainName) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String())
	}
	return nil
}

// Get the bytes for the message signer to sign on
func (msg RegisterDomainMessage) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

type ValidateDomainMessage struct {
	Sender     sdk.AccAddress
	DomainName string
	Owner      sdk.AccAddress
}

var _ sdk.Msg = ValidateDomainMessage{}

func NewValidateDomainMessage(sender sdk.AccAddress, domainName string, owner sdk.AccAddress) ValidateDomainMessage {
	return ValidateDomainMessage{
		Sender:     sender,
		DomainName: domainName,
		Owner:      owner,
	}
}

// enforce the msg type at compile time

// nolint
func (msg ValidateDomainMessage) Type() string                            { return "validate" }
func (msg ValidateDomainMessage) Get(key interface{}) (value interface{}) { return nil }
func (msg ValidateDomainMessage) GetSigners() []sdk.AccAddress            { return []sdk.AccAddress{msg.Sender} }
func (msg ValidateDomainMessage) String() string {
	return fmt.Sprintf("ValidateDomainMessage{Sender: %v, DomainName: %v, Owner: %v}", msg.Sender, msg.DomainName, msg.Owner)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg ValidateDomainMessage) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 || len(msg.DomainName) == 0 || len(msg.Owner) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String())
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

type ClaimRewardMessage struct {
	Sender     sdk.AccAddress
	DomainName string
	Fee        uint64
}

var _ sdk.Msg = ClaimRewardMessage{}

func NewClaimRewardMessage(sender sdk.AccAddress, domainName string, fee uint64) ClaimRewardMessage {
	return ClaimRewardMessage{
		Sender:     sender,
		DomainName: domainName,
		Fee:        fee,
	}
}

// enforce the msg type at compile time

// nolint
func (msg ClaimRewardMessage) Type() string                            { return "claim" }
func (msg ClaimRewardMessage) Get(key interface{}) (value interface{}) { return nil }
func (msg ClaimRewardMessage) GetSigners() []sdk.AccAddress            { return []sdk.AccAddress{msg.Sender} }
func (msg ClaimRewardMessage) String() string {
	return fmt.Sprintf("ClaimRewardMessage{Sender: %v, DomainName: %v}", msg.Sender, msg.DomainName)
}

// Validate Basic is used to quickly disqualify obviously invalid messages quickly
func (msg ClaimRewardMessage) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 || len(msg.DomainName) == 0 {
		return sdk.ErrUnknownAddress(msg.Sender.String())
	}
	return nil
}

// Get the bytes for the message signer to sign on
func (msg ClaimRewardMessage) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}
