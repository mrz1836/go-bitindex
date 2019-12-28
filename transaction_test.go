package bitindex

import "testing"

// TestClient_GetTransaction tests the GetTransaction()
func TestClient_GetTransaction(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var transaction *Transaction
	tx := "29d16d486955c1b1490f37d7ebc75a717ae2a61ca73b59c21fb6d586103e1cc7"
	transaction, err = client.GetTransaction(tx)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(transaction.ErrorMessage) > 0 {
		t.Fatal("error occurred", transaction.ErrorMessage, transaction.ErrorName, transaction.Errors)
	}

	if transaction.Hash != tx {
		t.Fatal("failed to get the hash", tx, transaction.Hash)
	}

	// Cant test for sats or balance, might change!
}

// TestClient_GetTransactionRaw tests the GetTransactionRaw()
func TestClient_GetTransactionRaw(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var rawTx *TransactionRaw
	tx := "29d16d486955c1b1490f37d7ebc75a717ae2a61ca73b59c21fb6d586103e1cc7"
	rawTx, err = client.GetTransactionRaw(tx)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(rawTx.ErrorMessage) > 0 {
		t.Fatal("error occurred", rawTx.ErrorMessage, rawTx.ErrorName, rawTx.Errors)
	}

	if len(rawTx.RawTx) == 0 {
		t.Fatal("failed to get the raw tx", tx, rawTx.RawTx)
	}

	// Cant test for sats or balance, might change!
}

// TestClient_SendTransaction tests the SendTransaction()
func TestClient_SendTransaction(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	//var resp *SendTransactionResponse
	rawTx := "0100000001d1bda0bde67183817b21af863adaa31fda8cafcf2083ca1eaba3054496cbde10010000006a47304402205fddd6abab6b8e94f36bfec51ba2e1f3a91b5327efa88264b5530d0c86538723022010e51693e3d52347d4d2ff142b85b460d3953e625d1e062a5fa2569623fb0ea94121029df3723daceb1fef64fa0558371bc48cc3a7a8e35d8e05b87137dc129a9d4598ffffffff0115d40000000000001976a91459cc95a8cde59ceda718dbf70e612dba4034552688ac00000000"
	_, err = client.SendTransaction(rawTx)
	if err == nil {
		t.Fatal("error should have occurred")
	}
}
