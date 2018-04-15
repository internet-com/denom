package denom

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Cool module reserves error 400-499 lawl
	CodeInvalidDenomRequest sdk.CodeType = 400
	CodeParameterMissing    sdk.CodeType = 401
)

// ErrIncorrectCoolAnswer - Error returned upon an incorrect guess
func ErrInvalidRequest(answer string) sdk.Error {
	return sdk.NewError(CodeInvalidDenomRequest, fmt.Sprintf("Invalid Message Request: %v", answer))
}

func ErrParameterMissing(parameter string) sdk.Error {
	return sdk.NewError(CodeParameterMissing, fmt.Sprintf("Parameter %v is missing.", parameter))
}
