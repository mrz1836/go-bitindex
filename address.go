package bitindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AddressInfo this endpoint retrieves various address info.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
func (c *Client) AddressInfo(address string) (addressInfo *AddressInfo, err error) {

	var resp string
	// /api/v3/network/addr/address
	resp, err = c.Request("addr/"+address, http.MethodGet, nil)
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
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
func (c *Client) AddressUnspentTransactions(address string) (transactions UnspentTransactions, err error) {

	var resp string
	// /api/v3/network/addr/address/utxo
	resp, err = c.Request("addr/"+address+"/utxo", http.MethodGet, nil)
	if err != nil {
		return
	}

	// Error?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", resp)
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
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
func (c *Client) GetTransactions(transactionRequest *GetTransactionsRequest) (response *GetTransactionsResponse, err error) {

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
	resp, err = c.Request("addrs/txs", http.MethodPost, data)
	if err != nil {
		return
	}

	response = new(GetTransactionsResponse)
	if err = json.Unmarshal([]byte(resp), &response); err != nil {
		return
	}
	return
}

// GetUnspentTransactions this endpoint retrieves list of unspent transactions.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Address
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
	resp, err = c.Request(endpoint, http.MethodPost, data)
	if err != nil {
		return
	}

	// Error?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", resp)
		return
	}

	transactions = *new(UnspentTransactions)
	if err = json.Unmarshal([]byte(resp), &transactions); err != nil {
		return
	}
	return
}
