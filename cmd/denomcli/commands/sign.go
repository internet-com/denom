package commands

import (
	"fmt"

	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/spf13/cobra"
)

func SignCommand() *cobra.Command {
	signCmd := &cobra.Command{
		Use:   "sign [data] [account]",
		Short: "Sign data using the private key of the specified account",
		Long:  `Sign data using the private key of the specified account`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			data := args[0]
			account := args[1]
			keybase, err := keys.GetKeyBase()
			if err != nil {
				return err
			}
			passphrase, err := ctx.GetPassphraseFromStdin(account)
			if err != nil {
				return err
			}
			signature, _, err := keybase.Sign(account, passphrase, []byte(data))
			if err != nil {
				return err
			}
			base64Str := base64.StdEncoding.EncodeToString(signature.Bytes())
			fmt.Println(base64Str)
			return nil
		},
	}
	return signCmd
}
