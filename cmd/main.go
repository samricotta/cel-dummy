package main

import (
	"github.com/spf13/cobra"
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
	
	submitCmd = &cobra.Command{
		Use:   "submit [namespace_id] [data] [gas_limit]",
		Short: "Submit a PayForData transaction to the celestia data availability network",
		Args:  cobra.ExactArgs(3),
		Run:   submit,
	}

	retrieveCmd = &cobra.Command{
		Use:   "retrieve [namespace_id] [block_height]",
		Short: "Retrieve namespaced data from the celestia data availability network",
		Args:  cobra.ExactArgs(2),
		Run:   retrieve,
	}
)
