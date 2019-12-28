package bitindex

import "testing"

// TestClient_ChainInfo tests the ChainInfo()
func TestClient_ChainInfo(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *ChainInfoResponse
	resp, err = client.ChainInfo()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Info.Difficulty == 0 {
		t.Fatal("we should have the Difficulty ", resp.Info.Difficulty)
	}
}

// TestClient_ChainDifficulty tests the ChainDifficulty()
func TestClient_ChainDifficulty(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *ChainDifficultyResponse
	resp, err = client.ChainDifficulty()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if resp.Difficulty == 0 {
		t.Fatal("we should have the Difficulty ", resp.Difficulty)
	}
}

// TestClient_ChainBestBlockHash tests the ChainBestBlockHash()
func TestClient_ChainBestBlockHash(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *ChainBestBlockHashResponse
	resp, err = client.ChainBestBlockHash()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.BestBlockHash) == 0 {
		t.Fatal("we should have a hash ", resp.BestBlockHash)
	}
}

// TestClient_ChainLastBlockHash tests the ChainLastBlockHash()
func TestClient_ChainLastBlockHash(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *ChainLastBlockHashResponse
	resp, err = client.ChainLastBlockHash()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.LastBlockHash) == 0 {
		t.Fatal("we should have a hash ", resp.LastBlockHash)
	}

	if len(resp.SyncTipHash) == 0 {
		t.Fatal("we should have a tip ", resp.SyncTipHash)
	}
}
