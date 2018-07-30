package commands

import (
	"encoding/base64"
	"errors"
	"fmt"

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
			info, err := keybase.Get(account)
			if err != nil {
				return err
			}
			// Verifying again to handle private key leaks due to the nature of ED25519
			// More context here: https://github.com/jedisct1/libsodium/issues/170
			pubKey := info.GetPubKey()
			pubKeyBytes := pubKey.Bytes()
			if pubKey.VerifyBytes([]byte(data), signature) {
				base64Sig := base64.StdEncoding.EncodeToString(signature.Bytes())
				base64PubKey := base64.StdEncoding.EncodeToString(pubKeyBytes)
				//fmt.Println("denomsig: " + base64Sig)
				//fmt.Println("denompubkey: " + base64PubKey)
				fmt.Printf("denomverify: key=%s;sig=%s\n", base64PubKey, base64Sig)
			} else {
				return errors.New("Unable to sign data")
			}
			return nil
		},
	}
	return signCmd
}
