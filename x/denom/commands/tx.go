package commands

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/builder"
	"github.com/cosmos/cosmos-sdk/wire"

	"github.com/svaishnavy/denom/x/denom"
)

// take the coolness quiz transaction
func SetDomainForSaleCommand(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "for_sale [domain] [fee]",
		Short: "set your domain for sale",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args[0]) == 0 {
				return errors.New("You must provide an answer")
			}

			// get the from address from the name flag
			from, err := builder.GetFromAddress()
			if err != nil {
				return err
			}

			fee, err := strconv.ParseUint(args[1], 10, 64)
			// create the message
			if err != nil {
				return err
			}
			msg := denom.NewSetDomainForSaleMessage(from, args[0], fee)
			fmt.Printf(msg.String())

			// get account name
			name := viper.GetString(client.FlagName)

			// build and sign the transaction, then broadcast to Tendermint
			res, err := builder.SignBuildBroadcast(name, msg, cdc)
			if err != nil {
				return err
			}

			fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
			return nil
		},
	}
}

// take the coolness quiz transaction
func RegisterDomainCommand(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "register_domain [domain] [fee]",
		Short: "Register domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args[0]) == 0 {
				return errors.New("You must provide an answer")
			}

			// get the from address from the name flag
			from, err := builder.GetFromAddress()
			if err != nil {
				return err
			}

			fee, err := strconv.ParseUint(args[1], 10, 64)
			// create the message
			if err != nil {
				return err
			}
			msg := denom.NewRegisterDomainMessage(from, args[0], fee)

			// get account name
			name := viper.GetString(client.FlagName)

			// build and sign the transaction, then broadcast to Tendermint
			res, err := builder.SignBuildBroadcast(name, msg, cdc)
			if err != nil {
				return err
			}

			fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
			return nil
		},
	}
}

// set a new cool trend transaction
func ValidateDomainCommand(cdc *wire.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "validate_domain [domain] [fee]",
		Short: "Validate domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args[0]) == 0 {
				return errors.New("You must provide an answer")
			}

			// get the from address from the name flag
			from, err := builder.GetFromAddress()
			if err != nil {
				return err
			}

			fee, err := strconv.ParseUint(args[1], 10, 64)
			// create the message
			if err != nil {
				return err
			}
			msg := denom.NewValidateDomainMessage(from, args[0], fee)

			// get account name
			name := viper.GetString(client.FlagName)

			// build and sign the transaction, then broadcast to Tendermint
			res, err := builder.SignBuildBroadcast(name, msg, cdc)
			if err != nil {
				return err
			}

			fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
			return nil
		},
	}
}
