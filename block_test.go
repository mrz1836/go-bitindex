package bitindex

import "testing"

// TestClient_GetBlockHashByHeight tests the GetBlockHashByHeight()
func TestClient_GetBlockHashByHeight(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *BlockHashByHeightResponse
	height := 15000
	resp, err = client.GetBlockHashByHeight(int64(height))
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.BlockHash) == 0 {
		t.Fatal("we should have the hash", resp.BlockHash)
	}
}

// TestClient_GetBlockHeader tests the GetBlockHeader()
func TestClient_GetBlockHeader(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *BlockHeaderResponse
	hash := "000000005c3548bda7c3d3e2fb2db21f68b87f47dd8b7ed095dfeecee76038b7"
	resp, err = client.GetBlockHeader(hash)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Hash != hash {
		t.Fatal("we should have the hash", resp.Hash, hash)
	}
}

// TestClient_GetBlock tests the GetBlock()
func TestClient_GetBlock(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *BlockResponse
	hash := "000000005c3548bda7c3d3e2fb2db21f68b87f47dd8b7ed095dfeecee76038b7"
	resp, err = client.GetBlock(hash)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Hash != hash {
		t.Fatal("we should have the hash", resp.Hash, hash)
	}

	if len(resp.Tx) == 0 {
		t.Fatal("we should have some transactions", resp.Tx, hash)
	}
}

// TestClient_GetBlockRaw tests the GetBlockRaw()
func TestClient_GetBlockRaw(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *BlockRawResponse
	hash := "000000005c3548bda7c3d3e2fb2db21f68b87f47dd8b7ed095dfeecee76038b7"
	resp, err = client.GetBlockRaw(hash)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.RawBlock) == 0 {
		t.Fatal("we should have the raw block", resp.RawBlock, hash)
	}

	if len(resp.ErrorMessage) > 0 {
		t.Fatal("we should not have an error", resp.ErrorMessage, hash)
	}
}
