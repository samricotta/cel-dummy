package retriever

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Retriever contains all necessary components in order to
// retrieve a message from the Celestia Data Availability Network
// by a given namespaceID and given height.
type Retriever struct {
	endpoint string
}

// namespacedDataResponse represents the structure of a
// /namespaced_data response.
type namespacedDataResponse struct {
	Data   [][]byte `json:"data"`
	Height uint64   `json:"height"`
}

func NewRetriever(endpoint string) *Retriever {
	return &Retriever{
		endpoint: endpoint,
	}
}

// Retrieve makes an http GET request to the Retriever's endpoint and
// returns the namespaced data as a hexadecimal string and an error.
func (r *Retriever) Retrieve(namespaceID string, blockHeight int64) (string, error) {
	// 1. create an http.GET request to the r.endpoint at the endpoint `/namespaced_data/{namespaceID}/height/{height}`
	resp, err := http.Get(fmt.Sprintf("http://%s/namespaced_data/%s/height/%d", r.endpoint, namespaceID, blockHeight))
	if err != nil {
		return "", err
	}

	// 2. ioutil.Readall(response.Body)
	//Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 3. json unmarshal the response into namespacedDataResponse
	namespacedData := new(namespacedDataResponse)
	err = json.Unmarshal(body, namespacedData)
	if err != nil {
		return "", err
	}

	// 5. turn bytes into string
	joinedArray := bytes.Join(namespacedData.Data, []byte{})
	dataString := string(joinedArray)

	// 6. return that hexadecimal and nil
	return dataString, nil
}
