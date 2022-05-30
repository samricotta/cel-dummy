package dummy

import (
	"github.com/samricotta/cel-dummy/keyring"
	"github.com/samricotta/cel-dummy/retriever"
	"github.com/samricotta/cel-dummy/submitter"
)

// Dummy stores all necessary components in order to submit and
// retrieve messages from the Celestia Data Availability network.
type Dummy struct {
	// submitter
	sub *submitter.Submitter

	// retriever
	ret *retriever.Retriever

	// endpoint info for celestia-node
	endpoint string

	// any other values you need to access
}

func NewDummy(cfg Config) (*Dummy, error) {
	signer, err := keyring.NewKeyringSigner(cfg.KeystorePath)
	if err != nil {
		return nil, err
	}

	return &Dummy{
		sub:      submitter.NewSubmitter(signer, cfg.AppEndpoint),
		ret:      retriever.NewRetriever(cfg.NodeEndpoint),
		endpoint: cfg.AppEndpoint,
	}, nil
}

func (d *Dummy) Submit(namespaceID []byte, data []byte, gasLim uint64) (int64, error) {
	// accesses the `submitter`'s `Submit()` function
	return d.sub.Submit(namespaceID, data, gasLim)
}

func (d *Dummy) Retrieve(namespaceID string, blockHeight int64) (string, error) {
	// accesses the `retriever`'s `Retrieve()` function
	return d.ret.Retrieve(namespaceID, blockHeight)
}
