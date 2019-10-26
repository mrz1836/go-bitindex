package bitindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetXpubNextAddress this endpoint that gets the next address for a xpub and reserve if given.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Xpub
func (c *Client) GetXpubNextAddress(xPub string, reserveTimeSeconds int) (addresses XpubAddresses, err error) {

	endpoint := "xpub/" + xPub + "/addrs/next"
	if reserveTimeSeconds > 0 {
		endpoint = fmt.Sprintf("%s?reserveTime=%d", endpoint, reserveTimeSeconds)
	}

	// Create the request
	var resp string
	// /api/v3/network/xpub/xpub/addrs/next
	resp, err = c.Request(endpoint, http.MethodGet, nil)
	if err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		var apiError APIInternalError
		if err = json.Unmarshal([]byte(resp), &apiError); err != nil {
			return
		}
		err = fmt.Errorf("error: %s", apiError.ErrorMessage)
		return
	}

	// Process the response
	addresses = *new(XpubAddresses)
	if err = json.Unmarshal([]byte(resp), &addresses); err != nil {
		return
	}
	return
}

//https://api.bitindex.network/api/v3/network/xpub/xpub/addrs?offset=&limit=1000&order=desc&address=

// GetXpubAddresses this endpoint will return addresses for an xpub given the parameters.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Xpub
func (c *Client) GetXpubAddresses(xPub string, offset, limit int, order, filterByAddress string) (addresses XpubAddresses, err error) {

	endpoint := "xpub/" + xPub + "/addrs"

	// Set the params
	params := url.Values{}
	if offset > 0 {
		params.Add("offset", strconv.Itoa(offset))
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(limit))
	}
	if len(order) > 0 {
		params.Add("order", strings.ToLower(order))
	}
	if len(filterByAddress) > 0 {
		params.Add("address", strings.TrimSpace(filterByAddress))
	}

	// Set the dynamic query string
	queryString := params.Encode()
	if len(queryString) > 0 {
		endpoint = endpoint + "?" + queryString
	}

	// Create the request
	var resp string
	// /api/v3/network/xpub/xpub/addrs
	resp, err = c.Request(endpoint, http.MethodGet, nil)
	if err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		var apiError APIInternalError
		if err = json.Unmarshal([]byte(resp), &apiError); err != nil {
			return
		}
		err = fmt.Errorf("error: %s", apiError.ErrorMessage)
		return
	}

	// Process the response
	addresses = *new(XpubAddresses)
	if err = json.Unmarshal([]byte(resp), &addresses); err != nil {
		return
	}
	return
}

// GetXpubBalance this endpoint that gets the total balance for the xpub.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Xpub
func (c *Client) GetXpubBalance(xPub string) (balance *XpubBalance, err error) {

	// Create the request
	var resp string
	// /api/v3/network/xpub/xpub/status
	resp, err = c.Request("xpub/"+xPub+"/status", http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	balance = new(XpubBalance)
	if err = json.Unmarshal([]byte(resp), balance); err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", balance.ErrorMessage)
		return
	}
	return
}

// GetXpubUnspentTransactions this endpoint retrieves list of unspent transactions for a xpub address.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Xpub
func (c *Client) GetXpubUnspentTransactions(xPub, sort string) (transactions UnspentTransactions, err error) {

	endpoint := "xpub/" + xPub + "/utxo"
	if len(sort) > 0 {
		endpoint = endpoint + "?sort=" + sort
	}

	// Create the request
	var resp string
	// /api/v3/network/xpub/xpub/utxo
	resp, err = c.Request(endpoint, http.MethodGet, nil)
	if err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		var apiError APIInternalError
		if err = json.Unmarshal([]byte(resp), &apiError); err != nil {
			return
		}
		err = fmt.Errorf("error: %s", apiError.ErrorMessage)
		return
	}

	// Process the response
	transactions = *new(UnspentTransactions)
	if err = json.Unmarshal([]byte(resp), &transactions); err != nil {
		return
	}
	return
}

// GetXpubTransactions this endpoint that gets the the history of transactions for the xpub.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Xpub
func (c *Client) GetXpubTransactions(xPub string) (transactions XpubAddresses, err error) {

	// Create the request
	var resp string
	// /api/v3/network/xpub/xpub/txs
	resp, err = c.Request("xpub/"+xPub+"/txs", http.MethodGet, nil)
	if err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		var apiError APIInternalError
		if err = json.Unmarshal([]byte(resp), &apiError); err != nil {
			return
		}
		err = fmt.Errorf("error: %s", apiError.ErrorMessage)
		return
	}

	// Process the response
	transactions = *new(XpubAddresses)
	if err = json.Unmarshal([]byte(resp), &transactions); err != nil {
		return
	}
	return
}
