package bitindex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTransaction this endpoint retrieves the transaction info.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Transactions
func (c *Client) GetTransaction(txID string) (transaction *Transaction, err error) {

	var resp string
	// /api/v3/network/tx/txid
	resp, err = c.Request("tx/"+txID, http.MethodGet, nil)
	if err != nil {
		return
	}

	transaction = new(Transaction)
	if err = json.Unmarshal([]byte(resp), transaction); err != nil {
		return
	}
	return
}

// GetTransactionRaw this endpoint retrieves the transaction in raw format.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Transactions
func (c *Client) GetTransactionRaw(txID string) (rawTx *TransactionRaw, err error) {

	var resp string
	// /api/v3/network/rawtx/txid
	resp, err = c.Request("rawtx/"+txID, http.MethodGet, nil)
	if err != nil {
		return
	}

	rawTx = new(TransactionRaw)
	if err = json.Unmarshal([]byte(resp), rawTx); err != nil {
		return
	}
	return
}

// SendTransaction this endpoint broadcasts a raw transaction to the network.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Transactions
func (c *Client) SendTransaction(rawTx string) (txID *SendTransactionResponse, err error) {

	var resp string
	// /api/v3/network/tx/send
	resp, err = c.Request("tx/send", http.MethodPost, []byte(fmt.Sprintf(`{"rawtx":"%s"}`, rawTx)))
	if err != nil {
		return
	}

	txID = new(SendTransactionResponse)
	if err = json.Unmarshal([]byte(resp), txID); err != nil {
		return
	}
	return
}
