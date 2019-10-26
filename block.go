package bitindex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetBlockHashByHeight this endpoint retrieves the block hash by height.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Block
func (c *Client) GetBlockHashByHeight(height int64) (blockHash *BlockHashByHeightResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/block-index/height
	resp, err = c.Request(fmt.Sprintf("block-index/%d", height), http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	blockHash = new(BlockHashByHeightResponse)
	if err = json.Unmarshal([]byte(resp), blockHash); err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", blockHash.ErrorMessage)
		return
	}

	return
}

// GetBlockHeader this endpoint retrieves the block header by hash.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Block
func (c *Client) GetBlockHeader(hash string) (blockHeader *BlockHeaderResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/blockheader/hash
	resp, err = c.Request(fmt.Sprintf("blockheader/%s", hash), http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	blockHeader = new(BlockHeaderResponse)
	if err = json.Unmarshal([]byte(resp), blockHeader); err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", blockHeader.ErrorMessage)
		return
	}
	return
}

// GetBlock this endpoint retrieves the block by hash.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Block
func (c *Client) GetBlock(hash string) (block *BlockResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/block/hash
	resp, err = c.Request(fmt.Sprintf("block/%s", hash), http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	block = new(BlockResponse)
	if err = json.Unmarshal([]byte(resp), block); err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", block.ErrorMessage)
		return
	}
	return
}

// GetBlockRaw this endpoint retrieves the raw block by hash.
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Block
func (c *Client) GetBlockRaw(hash string) (rawBlock *BlockRawResponse, err error) {

	// Create the request
	var resp string
	// /api/v3/network/rawblock/hash
	resp, err = c.Request(fmt.Sprintf("rawblock/%s", hash), http.MethodGet, nil)
	if err != nil {
		return
	}

	// Process the response
	rawBlock = new(BlockRawResponse)
	if err = json.Unmarshal([]byte(resp), rawBlock); err != nil {
		return
	}

	// Error from request?
	if c.LastRequest.StatusCode != http.StatusOK {
		err = fmt.Errorf("error: %s", rawBlock.ErrorMessage)
		return
	}
	return
}
