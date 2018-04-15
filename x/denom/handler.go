package denom

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// This is just an example to demonstrate a functional custom module
// with full feature set functionality.
//
//  /$$$$$$$  /$$$$$$   /$$$$$$  /$$
// /$$_____/ /$$__  $$ /$$__  $$| $$
//| $$      | $$  \ $$| $$  \ $$| $$
//| $$      | $$  | $$| $$  | $$| $$
//|  $$$$$$$|  $$$$$$/|  $$$$$$/| $$$$$$$
// \_______/ \______/  \______/ |______/

// NewHandler returns a handler for "cool" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case ClaimDomainMessage:
			return handleClaimDomainMessage(ctx, k, msg)
		case SetDomainForSaleMessage:
			return handleSetDomainForSaleMessage(ctx, k, msg)
		case ValidateDomainMessage:
			return handleValidateDomainMessage(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleClaimDomainMessage(ctx sdk.Context, k Keeper, msg ClaimDomainMessage) sdk.Result {
	//k.setTrend(ctx, msg.Cool)
	k.Claim(ctx, msg.Sender, msg.DomainName)
	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

// Handle QuizMsg This is the engine of your module
func handleValidateDomainMessage(ctx sdk.Context, k Keeper, msg ValidateDomainMessage) sdk.Result {
	//k.setTrend(ctx, msg.Cool)
	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

// Handle QuizMsg This is the engine of your module
func handleSetDomainForSaleMessage(ctx sdk.Context, k Keeper, msg SetDomainForSaleMessage) sdk.Result {
	//correct := k.CheckTrend(ctx, msg.CoolAnswer)

	/*
		if !correct {
			return ErrInvalidRequest(msg.CoolAnswer).Result()
		}

		if ctx.IsCheckTx() {
			return sdk.Result{} // TODO
		}

		bonusCoins := sdk.Coins{{msg.CoolAnswer, 69}}

		_, err := k.ck.AddCoins(ctx, msg.Sender, bonusCoins)
		if err != nil {
			return err.Result()
		}
	*/

	fmt.Println("Got message: " + msg.String())

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}
