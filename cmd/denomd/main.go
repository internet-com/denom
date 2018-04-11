package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/tmlibs/cli"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/svaishnavy/denom/app"
)

// democoindCmd is the entry point for this binary
var (
	denomCmd = &cobra.Command{
		Use:   "denom",
		Short: "Denom Daemon (server)",
	}
)

// defaultOptions sets up the app_options for the
// default genesis file
func defaultOptions(args []string) (json.RawMessage, string, cmn.HexBytes, error) {
	addr, secret, err := server.GenerateCoinKey()
	if err != nil {
		return nil, "", nil, err
	}
	fmt.Println("Secret phrase to access coins:")
	fmt.Println(secret)

	opts := fmt.Sprintf(`{
      "accounts": [{
        "address": "%s",
        "coins": [
          {
            "denom": "DNOM",
            "amount": 9007199254740992
          }
        ]
      }]
    }`, addr)
	return json.RawMessage(opts), "", nil, nil
}

func generateApp(rootDir string, logger log.Logger) (abci.Application, error) {
	dbMain, err := dbm.NewGoLevelDB("denom", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbAcc, err := dbm.NewGoLevelDB("denom-acc", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbIBC, err := dbm.NewGoLevelDB("denom-ibc", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbStaking, err := dbm.NewGoLevelDB("denom-staking", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbs := map[string]dbm.DB{
		"main":    dbMain,
		"acc":     dbAcc,
		"ibc":     dbIBC,
		"staking": dbStaking,
	}
	bapp := app.NewDenomApp(logger, dbs)
	return bapp, nil
}

func main() {
	// TODO: set logger through CLI
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).
		With("module", "main")

	denomCmd.AddCommand(
		server.InitCmd(defaultOptions, logger),
		server.StartCmd(generateApp, logger),
		server.UnsafeResetAllCmd(logger),
		server.ShowNodeIdCmd(logger),
		server.ShowValidatorCmd(logger),
		version.VersionCmd,
	)

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/.denom")
	executor := cli.PrepareBaseCmd(denomCmd, "BC", rootDir)
	executor.Execute()
}
