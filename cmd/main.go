package main

import (
	"context"
	"github.com/celestiaorg/celestia-app/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"os"
)

// Commands
var (
	rootCmd = &cobra.Command{
		Use: "cel-dummy",
		Short: `
 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄               ▄▄▄▄▄▄▄▄▄▄   ▄         ▄  ▄▄       ▄▄  ▄▄       ▄▄  ▄         ▄
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌             ▐░░░░░░░░░░▌ ▐░▌       ▐░▌▐░░▌     ▐░░▌▐░░▌     ▐░░▌▐░▌       ▐░▌
▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌             ▐░█▀▀▀▀▀▀▀█░▌▐░▌       ▐░▌▐░▌░▌   ▐░▐░▌▐░▌░▌   ▐░▐░▌▐░▌       ▐░▌
▐░▌          ▐░▌          ▐░▌             ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌▐░▌ ▐░▌▐░▌▐░▌▐░▌ ▐░▌▐░▌▐░▌       ▐░▌
▐░▌          ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌ ▄▄▄▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌ ▐░▐░▌ ▐░▌▐░▌ ▐░▐░▌ ▐░▌▐░█▄▄▄▄▄▄▄█░▌
▐░▌          ▐░░░░░░░░░░░▌▐░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░▌       ▐░▌▐░▌  ▐░▌  ▐░▌▐░▌  ▐░▌  ▐░▌▐░░░░░░░░░░░▌
▐░▌          ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌ ▀▀▀▀▀▀▀▀▀▀▀ ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌   ▀   ▐░▌▐░▌   ▀   ▐░▌ ▀▀▀▀█░█▀▀▀▀
▐░▌          ▐░▌          ▐░▌             ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌
▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄    ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌   ▐░░░░░░░░░░▌ ▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌
 ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀     ▀▀▀▀▀▀▀▀▀▀   ▀▀▀▀▀▀▀▀▀▀▀  ▀         ▀  ▀         ▀       ▀


Use cel-dummy to submit and retrieve messages to/from the celestia data availability network`,
		Args: cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	// TODO fill out the submit command
	submitCmd = &cobra.Command{
		Use:   "submit [namespace_id] [data] [gas_limit]",
		Short: "Submit a PayForData transaction to the celestia data availability network",
		Args:  cobra.ExactArgs(3),
		Run:   submit,
	}

	// TODO fill out the retrieve command
	retrieveCmd = &cobra.Command{
		Use:   "retrieve [namespace_id] [block_height]",
		Short: "Retrieve namespaced data from the celestia data availability network",
		Args:  cobra.ExactArgs(2),
		Run:   retrieve,
	}
)

func init() {
	// This is necessary to ensure that the account addresses are correctly prefixed
	// as in the celestia application.
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	cfg.Seal()

	// Add submitter and retriever subcommands
	rootCmd.AddCommand(
		submitCmd,
		retrieveCmd,
	)

	rootCmd.SetHelpCommand(&cobra.Command{})
}

func main() {
	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		os.Exit(1)
	}
}
