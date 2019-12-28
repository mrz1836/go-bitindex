package bitindex

import "testing"

// TestClient_GetWebhookConfig tests the GetWebhookConfig()
func TestClient_GetWebhookConfig(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp *WebhookConfigResponse
	resp, err = client.GetWebhookConfig()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp.URL) == 0 {
		t.Fatal("we should have the url", resp.URL)
	}
}

// TestClient_GetMonitoredAddresses tests the GetMonitoredAddresses()
func TestClient_GetMonitoredAddresses(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var resp MonitoredAddresses
	resp, err = client.GetMonitoredAddresses()
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp) == 0 {
		t.Fatal("should have monitored addresses", resp)
	}
}

// TestClient_UpdateWebhookConfig tests the UpdateWebhookConfig()
func TestClient_UpdateWebhookConfig(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var config *WebhookConfigResponse

	req := new(WebhookUpdateConfig)
	req.Secret = "Top-Secret-New-Key2"
	req.Enabled = false
	//req.URL = "http://example-update-me-to-your.com/path/callback2"

	config, err = client.UpdateWebhookConfig(req)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(config.ID) == 0 {
		t.Fatal("we should at least have an id", config.ID)
	}

	if config.Secret != req.Secret {
		t.Fatal("secret should have updated", req.Secret, config.Secret)
	}
}

// TestClient_AddMonitoredAddresses tests the AddMonitoredAddresses()
func TestClient_AddMonitoredAddresses(t *testing.T) {
	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(testAPIKey, NetworkMain, nil)
	if err != nil {
		t.Fatal(err)
	}

	var req MonitoredAddresses
	var add MonitoredAddress
	add.Address = "1M6N389jhRi5DQgoQcNir2e2REpYeAYavD"
	req = append(req, add)

	var resp MonitoredAddresses
	resp, err = client.AddMonitoredAddresses(&req)
	if err != nil {
		t.Fatal("error occurred: " + err.Error())
	}

	if len(resp) == 0 {
		t.Fatal("should have monitored addresses", resp)
	}

	if resp[0].Address != add.Address {
		t.Fatal("address returned should be the one we added", resp[0].Address, add.Address)
	}
}
