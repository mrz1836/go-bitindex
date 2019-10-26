/*
Package bitindex is the unofficial golang implementation for the bitindex API

Example:

// Create a client
client, _ := bitindex.NewClient()
*/
package bitindex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

// Client holds client configuration settings
type Client struct {

	// HTTPClient carries out the POST operations
	HTTPClient heimdall.Client

	// Parameters contains the search parameters that are submitted with your query,
	// which may affect the data returned
	Parameters *RequestParameters

	// LastRequest is the raw information from the last request
	LastRequest *LastRequest
}

// RequestParameters holds options that can affect data returned by a request.
type RequestParameters struct {

	// UserAgent (optional for changing user agents)
	UserAgent string

	// Network is what this search should use IE: main
	Network NetworkType

	// APIKey is the api key to use
	APIKey string
}

// LastRequest is used to track what was submitted to the Request()
type LastRequest struct {

	// Method is either POST or GET
	Method string

	// PostData is the post data submitted if POST request
	PostData string

	// StatusCode is the last code from the request
	StatusCode int

	// URL is the url used for the request
	URL string
}

// NewClient creates a new client to submit queries with.
// Parameters values are set to the defaults defined by WhatsOnChain.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Authentication
func NewClient(apiKey string) (c *Client, err error) {

	// Make sure we have an API key
	if len(apiKey) == 0 {
		err = fmt.Errorf("missing required api key")
		return
	}

	// Create a client
	c = new(Client)

	// Create exponential backoff
	backOff := heimdall.NewExponentialBackoff(
		ConnectionInitialTimeout,
		ConnectionMaxTimeout,
		ConnectionExponentFactor,
		ConnectionMaximumJitterInterval,
	)

	// Create the http client
	//c.HTTPClient = new(http.Client) (@mrz this was the original HTTP client)
	c.HTTPClient = httpclient.NewClient(
		httpclient.WithHTTPTimeout(ConnectionWithHTTPTimeout),
		httpclient.WithRetrier(heimdall.NewRetrier(backOff)),
		httpclient.WithRetryCount(ConnectionRetryCount),
		httpclient.WithHTTPClient(&http.Client{
			Transport: ClientDefaultTransport,
		}),
	)

	// Create default parameters
	c.Parameters = new(RequestParameters)
	c.Parameters.APIKey = apiKey
	c.Parameters.Network = NetworkMain
	c.Parameters.UserAgent = DefaultUserAgent

	// Create a last request struct
	c.LastRequest = new(LastRequest)

	// Return the client
	return
}

// Request is a generic bitindex request wrapper that can be used without constraints
func (c *Client) Request(endpoint string, method string, payload []byte) (response string, err error) {

	// Set reader
	var bodyReader io.Reader

	// Add the network value
	endpoint = fmt.Sprintf("%s%s/%s", APIEndpoint, c.Parameters.Network, endpoint)

	// Switch on POST vs GET
	switch method {
	case "POST":
		{
			bodyReader = bytes.NewBuffer(payload)
		}
	}

	// Store for debugging purposes
	c.LastRequest.Method = method
	c.LastRequest.URL = endpoint

	// Start the request
	var request *http.Request
	if request, err = http.NewRequest(method, endpoint, bodyReader); err != nil {
		return
	}

	// Change the header (user agent is in case they block default Go user agents)
	request.Header.Set("User-Agent", c.Parameters.UserAgent)
	request.Header.Set("api_key", c.Parameters.APIKey)

	// Set the content type on POST
	if method == "POST" {
		request.Header.Set("Content-Type", "application/json")
	}

	// Fire the http request
	var resp *http.Response
	if resp, err = c.HTTPClient.Do(request); err != nil {
		return
	}

	// Close the response body
	defer func() {
		if bodyErr := resp.Body.Close(); bodyErr != nil {
			log.Printf("error closing response body: %s", bodyErr.Error())
		}
	}()

	// Save the status
	c.LastRequest.StatusCode = resp.StatusCode

	// Read the body
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// Parse the response
	response = string(body)

	// Done
	return
}

// AddressInfo this endpoint retrieves various address info.
//
// Form more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
func (c *Client) AddressInfo(address string) (addressInfo *AddressInfo, err error) {

	var resp string
	// /api/v3/network/addr/address
	resp, err = c.Request("addr/"+address, "GET", nil)
	if err != nil {
		return
	}

	addressInfo = new(AddressInfo)
	if err = json.Unmarshal([]byte(resp), addressInfo); err != nil {
		return
	}
	return
}

// AddressUnspentTransactions this endpoint retrieves list of UTXOs.
//
// Form more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
func (c *Client) AddressUnspentTransactions(address string) (transactions UnspentTransactions, err error) {

	var resp string
	// /api/v3/network/addr/address/utxo
	resp, err = c.Request("addr/"+address+"/utxo", "GET", nil)
	if err != nil {
		return
	}

	transactions = *new(UnspentTransactions)
	if err = json.Unmarshal([]byte(resp), &transactions); err != nil {
		return
	}
	return
}

// GetTransactions this endpoint retrieves list of transactions.
//
// Form more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
func (c *Client) GetTransactions(transactionRequest *GetTransactionsRequest) (response GetTransactionsResponse, err error) {

	// Got multiple addresses?
	if len(transactionRequest.Addresses) > 0 {
		transactionRequest.Address = strings.Join(transactionRequest.Addresses, ",")
	}

	var data []byte
	data, err = json.Marshal(transactionRequest)
	if err != nil {
		return
	}

	var resp string
	// /api/v3/network/addrs/txs
	resp, err = c.Request("addrs/txs", "POST", data)
	if err != nil {
		return
	}

	response = *new(GetTransactionsResponse)
	if err = json.Unmarshal([]byte(resp), &response); err != nil {
		return
	}
	return
}

// GetUnspentTransactions this endpoint retrieves list of unspent transactions.
//
// Form more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
func (c *Client) GetUnspentTransactions(transactionRequest *GetUnspentTransactionsRequest) (transactions UnspentTransactions, err error) {

	// Got multiple addresses?
	if len(transactionRequest.Addresses) > 0 {
		transactionRequest.Address = strings.Join(transactionRequest.Addresses, ",")
	}

	var data []byte
	data, err = json.Marshal(transactionRequest)
	if err != nil {
		return
	}

	// Do we have a sort
	endpoint := "addrs/utxo"
	if len(transactionRequest.Sort) > 0 {
		endpoint = fmt.Sprintf("%s?sort=%s", endpoint, transactionRequest.Sort)
	}

	var resp string
	// /api/v3/network/addrs/utxo
	resp, err = c.Request(endpoint, "POST", data)
	if err != nil {
		return
	}

	transactions = *new(UnspentTransactions)
	if err = json.Unmarshal([]byte(resp), &transactions); err != nil {
		return
	}
	return
}
