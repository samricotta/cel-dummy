package keyring

import (
	"fmt"
	"github.com/celestiaorg/celestia-app/app"
	"github.com/celestiaorg/celestia-app/app/encoding"
	"github.com/celestiaorg/celestia-app/x/payment/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"os"
)

/*
Construct a keyring signer here that is used by the Submitter

Copy keyring construction from
https://github.com/celestiaorg/celestia-node/blob/main/node/state/keyring.go
*/

func NewKeyringSigner(keystorePath string) (*types.KeyringSigner, error) {
	// create new keyring
	encConf := encoding.MakeEncodingConfig(app.ModuleEncodingRegisters...)
	ring, err := keyring.New(app.Name, keyring.BackendTest, keystorePath, os.Stdin, encConf.Codec)
	if err != nil {
		return nil, err
	}
	keys, err := ring.List()
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, fmt.Errorf("keys not found in keystore path: %s", keystorePath)
	}
	info := keys[0]

	// hardcoding the network "mamaki", construct the signer
	signer := types.NewKeyringSigner(ring, info.Name, "mamaki")
	return signer, nil
}
