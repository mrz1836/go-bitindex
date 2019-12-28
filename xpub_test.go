package bitindex

import "testing"

// TestClient_GetXpubNextAddress tests the GetXpubNextAddress()
func TestClient_GetXpubNextAddress(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var addresses XpubAddresses
	xPub := "xpub6AHA9hZDN11k2ijHMeS5QqHx2KP9aMBRhTDqANMnwVtdyw2TDYRmF8PjpvwUFcL1Et8Hj59S3gTSMcUQ5gAqTz3Wd8EsMTmF3DChhqPQBnU"
	seconds := 0 // testing with NO reserve time
	addresses, err = client.GetXpubNextAddress(xPub, seconds)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(addresses) == 0 {
		t.Fatal("we should have the some addresses", addresses, xPub)
	}

	seconds = 60 // testing with reserve time
	addresses, err = client.GetXpubNextAddress(xPub, seconds)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(addresses) == 0 {
		t.Fatal("we should have the some addresses", addresses, xPub)
	}
}

// TestClient_GetXpubAddresses tests the GetXpubAddresses()
func TestClient_GetXpubAddresses(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var addresses XpubAddresses
	xPub := "xpub6AHA9hZDN11k2ijHMeS5QqHx2KP9aMBRhTDqANMnwVtdyw2TDYRmF8PjpvwUFcL1Et8Hj59S3gTSMcUQ5gAqTz3Wd8EsMTmF3DChhqPQBnU"
	offset := 0           // testing with NO offset
	limit := 0            // testing with NO limit
	order := ""           // testing with NO order
	filterByAddress := "" // testing with NO filter
	addresses, err = client.GetXpubAddresses(xPub, offset, limit, order, filterByAddress)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(addresses) == 0 {
		t.Fatal("we should have the some addresses", addresses, xPub)
	}

	// Test Limit
	limit = 5 // testing with limit
	addresses, err = client.GetXpubAddresses(xPub, offset, limit, order, filterByAddress)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(addresses) != 5 {
		t.Fatal("we should have 5 addresses", addresses, xPub, limit)
	}
}

// TestClient_GetXpubBalance tests the GetXpubBalance()
func TestClient_GetXpubBalance(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var balance *XpubBalance
	xPub := "xpub6AHA9hZDN11k2ijHMeS5QqHx2KP9aMBRhTDqANMnwVtdyw2TDYRmF8PjpvwUFcL1Et8Hj59S3gTSMcUQ5gAqTz3Wd8EsMTmF3DChhqPQBnU"

	balance, err = client.GetXpubBalance(xPub)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if balance.ErrorMessage != "" {
		t.Fatal("error:", balance.ErrorMessage)
	}

	// Cannot safely test the balance
}

// TestClient_GetXpubUnspentTransactions tests the GetXpubUnspentTransactions()
func TestClient_GetXpubUnspentTransactions(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var transactions UnspentTransactions
	xPub := "xpub6AHA9hZDN11k2ijHMeS5QqHx2KP9aMBRhTDqANMnwVtdyw2TDYRmF8PjpvwUFcL1Et8Hj59S3gTSMcUQ5gAqTz3Wd8EsMTmF3DChhqPQBnU"
	sort := "" // not testing for sort
	transactions, err = client.GetXpubUnspentTransactions(xPub, sort)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(transactions) == 0 {
		//t.Fatal("we should have some transactions!")
		t.Log("found transactions!")
	} else {
		t.Log("no transactions found")
	}

	// Cant test for sats or balance, might change!
}

// TestClient_GetXpubTransactions tests the GetXpubTransactions()
func TestClient_GetXpubTransactions(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var transactions XpubAddresses
	xPub := "xpub6AHA9hZDN11k2ijHMeS5QqHx2KP9aMBRhTDqANMnwVtdyw2TDYRmF8PjpvwUFcL1Et8Hj59S3gTSMcUQ5gAqTz3Wd8EsMTmF3DChhqPQBnU"
	transactions, err = client.GetXpubTransactions(xPub)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(transactions) == 0 {
		t.Fatal("we should have the some transactions", transactions, xPub)
	}
}
