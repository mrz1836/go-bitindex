package bitindex

import "testing"

// TestClient_AddressInfo tests the AddressInfo()
func TestClient_AddressInfo(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	var resp *AddressInfo
	address := "16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA"
	resp, err = client.AddressInfo(address)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.ErrorMessage) > 0 {
		t.Fatal("error occurred", resp.ErrorMessage, resp.ErrorName, resp.Errors)
	}

	if resp.Address != address {
		t.Fatal("failed to get the address", address, resp.Address)
	}

	// Cant test for sats or balance, might change!
}

// TestClient_AddressUnspentTransactions tests the AddressUnspentTransactions()
func TestClient_AddressUnspentTransactions(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	var transactions UnspentTransactions
	address := "16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA"
	transactions, err = client.AddressUnspentTransactions(address)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(transactions) == 0 {
		t.Fatal("error occurred: missing transactions")
	}

	// Cant test for sats or balance, might change!
}

// TestClient_GetTransactions tests the GetTransactions()
func TestClient_GetTransactions(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	var resp *GetTransactionsResponse
	address := "16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA"

	req := new(GetTransactionsRequest)
	req.Address = address
	req.IncludeAsm = true
	req.IncludeHex = true

	resp, err = client.GetTransactions(req)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.ErrorMessage) > 0 {
		t.Fatal("error occurred", resp.ErrorMessage, resp.ErrorName, resp.Errors)
	}

	if len(resp.Items) == 0 {
		t.Fatal("we should have some items!")
	}

	// Cant test for sats or balance, might change!
}

// TestClient_GetTransactionsMultiple tests the GetTransactions()
func TestClient_GetTransactionsMultiple(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	var resp *GetTransactionsResponse
	address := "16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA"

	req := new(GetTransactionsRequest)
	req.Addresses = append(req.Addresses, address, "154d9jp3VZx9kZGtUqFvCWSrV94p5s6zdB")
	req.IncludeAsm = true
	req.IncludeHex = true

	resp, err = client.GetTransactions(req)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.ErrorMessage) > 0 {
		t.Fatal("error occurred", resp.ErrorMessage, resp.ErrorName, resp.Errors)
	}

	if len(resp.Items) == 0 {
		t.Fatal("we should have some items!")
	}

	// Cant test for sats or balance, might change!
}

// TestClient_GetUnspentTransactions tests the GetUnspentTransactions()
func TestClient_GetUnspentTransactions(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	var transactions UnspentTransactions
	address := "16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA"

	req := new(GetUnspentTransactionsRequest)
	req.Address = address
	req.Sort = "value:desc"

	transactions, err = client.GetUnspentTransactions(req)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(transactions) == 0 {
		t.Fatal("we should have some transactions!")
	}

	// Cant test for sats or balance, might change!
}
