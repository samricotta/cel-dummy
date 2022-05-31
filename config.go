package dummy

import "path/filepath"

// Config contains all configuration values for
// creating the Dummy.
type Config struct {
	// path to directory where the keys are stored
	KeystorePath string

	// AppEndpoint of celestia application instance
	AppEndpoint string

	// NodeEndpoint of celestia data availability node
	NodeEndpoint string
}

func DefaultConfig() Config {
	ksPath, err := filepath.Abs("./keys")
	if err != nil {
	}

	return Config{
		KeystorePath: ksPath,
		// Hardcoded endpoint of celestia app
		AppEndpoint: "46.101.42.40:9090",
		// Hardcoded endpoint of celestia light node
		NodeEndpoint: "127.0.0.1:26658",
	}
}
