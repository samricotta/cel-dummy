package submitter

import (
	"context"
	"errors"
	"fmt"
	"github.com/celestiaorg/celestia-app/x/payment"
	"github.com/celestiaorg/celestia-app/x/payment/types"
	"github.com/samricotta/cel-dummy/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Submitter contains all components necessary to be able
// to submit a message to the Celestia Data Availability Network.
type Submitter struct {
	signer   *types.KeyringSigner
	endpoint string
}

func NewSubmitter(signer *types.KeyringSigner, endpoint string) *Submitter {
	return &Submitter{
		signer:   signer,
		endpoint: endpoint,
	}
}

// Submit submits a PayForData transaction with the given parameters, and upon success,
// returns the block height at which the PFD transaction was included.
func (s *Submitter) Submit(namespaceID []byte, data []byte, gasLim uint64) (int64, error) {
	// sanity check the length of the namespaceID
	err := util.CheckNamespaceLength(namespaceID)
	if err != nil {
		fmt.Printf("%#v", err)
		return 0, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	// dial the gRPC endpoint
	client, err := grpc.DialContext(ctx, s.endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("%#v", client)
		return 0, err
	}

	resp, err := payment.SubmitPayForData(ctx, s.signer, client, namespaceID, data, gasLim)
	if err != nil {
		fmt.Printf("%#v", err)
		return 0, err
	}
	// if height is 0, an error occurred, return
	if resp.Height == 0 {
		fmt.Printf("%#v%v\n", resp)
		return 0, errors.New(resp.Info)
	}
	return resp.Height, nil
}
