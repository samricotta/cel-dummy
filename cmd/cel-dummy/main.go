package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/celestiaorg/celestia-app/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dummy "github.com/samricotta/cel-dummy"
	"github.com/spf13/cobra"
	"os"
	"strconv"
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

func submit(cmd *cobra.Command, args []string) {
	// parse arguments
	namespaceID, data, gasLim, err := parseSubmitArgs(args[0], args[1], args[2])
	if err != nil {
		fmt.Println("Err while parsing submit arguments: ", err)
		os.Exit(1)
	}
	// construct the dummy app
	dummy, err := dummy.NewDummy(dummy.DefaultConfig())
	if err != nil {
		fmt.Println("Err while constructing new dummy: ", err)
		os.Exit(1)
	}
	resp, err := dummy.Submit(namespaceID, data, gasLim)
	if err != nil {
		fmt.Printf("Err while submitting PayForData: %v\n", err)
		os.Exit(1)

	}
	fmt.Println("Message successfully submitted at block height: ", resp)
}

func retrieve(cmd *cobra.Command, args []string) {
	namespaceID, blockHeight, err := parseRetrieveArgs(args[0], args[1])
	if err != nil {
		fmt.Println("Err while parsing submit arguments: ", err)
		os.Exit(1)
	}
	fmt.Println(namespaceID, blockHeight)
	// construct the dummy app
	dummy, err := dummy.NewDummy(dummy.DefaultConfig())
	if err != nil {
		fmt.Println("Err while constructing new dummy: ", err)
		os.Exit(1)
	}
	resp, err := dummy.Retrieve(namespaceID, blockHeight)
	if err != nil {
		fmt.Println("Err while retrieving namespaced data: ", err)
		os.Exit(1)
	}
	fmt.Println("Message retrieved: ", resp)
}

func parseSubmitArgs(namespaceID string, data string, gasLim string) ([]byte, []byte, uint64, error) {
	nID := []byte(namespaceID)
	dataBytes := []byte(data)

	gasLimInt, err := strconv.Atoi(gasLim)
	if err != nil {
		return []byte{}, []byte{}, 0, err
	}

	return nID, dataBytes, uint64(gasLimInt), nil
}

// parseRetrieveArgs takes in the namespaceID and blockHeight as strings,
// and returns the namespaceID as a hexadecimal string, the block height as an
// int64, and an error.
func parseRetrieveArgs(namespaceID string, blockHeight string) (string, int64, error) {
	// TODO @sam
	// cast namespaceID to bytes
	// encode namespaceIDbytes to hex
	nID := hex.EncodeToString([]byte(namespaceID))

	height, err := strconv.Atoi(blockHeight)
	if err != nil {
		return "", 0, err
	}

	return nID, int64(height), nil
}
