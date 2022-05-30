import (
"context"
"errors"
"github.com/celestiaorg/celestia-app/x/payment"
"github.com/celestiaorg/celestia-app/x/payment/types"
"github.com/renaynay/sam-app/util"
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