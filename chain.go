package bitindex

import (
	"encoding/json"
	"net/http"
)

// ChainInfo this endpoint retrieves the current chain info.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#ChainInfo
func (c *Client) ChainInfo() (chainInfo *ChainInfoResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/status
	resp, err = c.Request("status?q=chainInfo", http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	chainInfo = new(ChainInfoResponse)
	if err = json.Unmarshal([]byte(resp), chainInfo); err != nil {
		return
	}
	return
}

// ChainDifficulty this endpoint retrieves the current chain difficulty.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#ChainInfo
func (c *Client) ChainDifficulty() (difficulty *ChainDifficultyResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/status
	resp, err = c.Request("status?q=getDifficulty", http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	difficulty = new(ChainDifficultyResponse)
	if err = json.Unmarshal([]byte(resp), difficulty); err != nil {
		return
	}
	return
}

// ChainBestBlockHash this endpoint retrieves the current best block hash
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#ChainInfo
func (c *Client) ChainBestBlockHash() (bestBlockHash *ChainBestBlockHashResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/status
	resp, err = c.Request("status?q=getBestBlockHash", http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	bestBlockHash = new(ChainBestBlockHashResponse)
	if err = json.Unmarshal([]byte(resp), bestBlockHash); err != nil {
		return
	}
	return
}

// ChainLastBlockHash this endpoint retrieves the last block hash
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#ChainInfo
func (c *Client) ChainLastBlockHash() (lastBlockHash *ChainLastBlockHashResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/status
	resp, err = c.Request("status?q=getLastBlockHash", http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	lastBlockHash = new(ChainLastBlockHashResponse)
	if err = json.Unmarshal([]byte(resp), lastBlockHash); err != nil {
		return
	}
	return
}
