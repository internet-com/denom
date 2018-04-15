package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/builder"
	"github.com/cosmos/cosmos-sdk/wire"

	"github.com/svaishnavy/denom/x/denom"
)

// take the coolness quiz transaction
func SetDomainForSaleCommand(cdc *wire.Codec) *cobra.Command {
	command := &cobra.Command{
		Use:   "sell",
		Short: "set your domain for sale",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainName := viper.GetString("domain")
			// get the from address from the name flag
			from, err := builder.GetFromAddress()
			if err != nil {
				return err
			}

			fee, err := strconv.ParseUint(viper.GetString("fee"), 10, 64)

			//fee, err := strconv.ParseUint(args[1], 10, 64)
			// create the message
			if err != nil {
				return err
			}
			msg := denom.NewSetDomainForSaleMessage(from, domainName, fee)
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
	command.Flags().StringP("domain", "d", "", "domain name")
	return command
}

// take the coolness quiz transaction
func ClaimDomainCommand(cdc *wire.Codec) *cobra.Command {
	command := &cobra.Command{
		Use:   "claim",
		Short: "Claim your domain",
		RunE: func(cmd *cobra.Command, args []string) error {

			domainName := viper.GetString("domain")
			if domainName == "" {
				return denom.ErrParameterMissing("domain")
			}
			// get the from address from the name flag
			from, err := builder.GetFromAddress()
			if err != nil {
				return err
			}

			fee, err := strconv.ParseUint(viper.GetString("fee"), 10, 64)

			//fee, err := strconv.ParseUint(args[1], 10, 64)
			// create the message
			if err != nil {
				return err
			}
			msg := denom.NewClaimDomainMessage(from, domainName, fee)

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
	command.Flags().StringP("domain", "d", "", "domain name")
	//command.Flags().Int64("fees", 0, "fee")
	return command
}

// set a new cool trend transaction
func ValidateDomainCommand(cdc *wire.Codec) *cobra.Command {
	command := &cobra.Command{
		Use:   "validate",
		Short: "Validate domain(only by validators)",
		RunE: func(cmd *cobra.Command, args []string) error {

			domainName := viper.GetString("domain")
			// get the from address from the name flag
			from, err := builder.GetFromAddress()
			if err != nil {
				return err
			}

			fee, err := strconv.ParseUint(viper.GetString("fee"), 10, 64)

			//fee, err := strconv.ParseUint(args[1], 10, 64)
			// create the message
			if err != nil {
				return err
			}
			msg := denom.NewValidateDomainMessage(from, domainName, fee)

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
	command.Flags().StringP("domain", "d", "", "domain name")
	return command
}
