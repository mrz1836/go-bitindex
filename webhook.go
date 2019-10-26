package bitindex

import (
	"encoding/json"
	"net/http"
)

// Sample Incoming Webhook:
//
// Notify at the webhook endpoint callback url when an address or child address of an xpub receives a paymentCallbacks
// for addresses are received at your configured URL.Note: You can receive up to multiple callbacks in any order.
// Make sure to check the 'confirmations' parameter and always use the highest 'confirmations' your application has seen before.
// It is possible that old webhooks are in transit with a lower 'confirmations' than what you have received before.
// Note: You should be able to rely on payments of 3 confirmations.
// Always check with the > and < operators since it cannot be guaranteed that you will receive a webhook
// with exactly 3 confirmations (it could be 4, 5 or more).
/*
{
    txid: 'e9865ab744ef236f0f436455a439263a53d9708f5eca66625dccb85cf1ff5947',
    address: '1M6N389jhRi5DQgoQcNir2e2REpYeAYavD',
    xpub: 'xpub6CYu...',    // Xpub will be present if address is associated with an xpub
    path: '1/0',            // Path is set if xpub is present
    satoshis: 1273,
    confirmations: 3,
    vout: 0,
    secret: "secret123key", // Set this secret key above and then compare in your code
    network: "main",        // only main supported for now.
}
*/

// GetWebhookConfig this endpoint retrieves the configuration for the existing webhook
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Webhook
func (c *Client) GetWebhookConfig() (config *WebhookConfigResponse, err error) {

	var resp string
	// /api/v3/network/webhook/endpoint
	resp, err = c.Request("webhook/endpoint", http.MethodGet, nil)
	if err != nil {
		return
	}

	config = new(WebhookConfigResponse)
	if err = json.Unmarshal([]byte(resp), config); err != nil {
		return
	}
	return
}

// UpdateWebhookConfig this endpoint updates the configuration for the existing webhook
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Webhook
func (c *Client) UpdateWebhookConfig(updateConfig *WebhookUpdateConfig) (config *WebhookConfigResponse, err error) {

	var data []byte
	data, err = json.Marshal(updateConfig)
	if err != nil {
		return
	}

	var resp string
	// /api/v3/network/webhook/endpoint
	resp, err = c.Request("webhook/endpoint", http.MethodPut, data)
	if err != nil {
		return
	}

	config = new(WebhookConfigResponse)
	if err = json.Unmarshal([]byte(resp), config); err != nil {
		return
	}
	return
}

// GetMonitoredAddresses this endpoint retrieves all the addresses being monitored by that API key
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Webhook
func (c *Client) GetMonitoredAddresses() (addresses MonitoredAddresses, err error) {

	var resp string
	// /api/v3/network/webhook/monitored_addrs
	resp, err = c.Request("webhook/monitored_addrs", http.MethodGet, nil)
	if err != nil {
		return
	}

	addresses = *new(MonitoredAddresses)
	if err = json.Unmarshal([]byte(resp), &addresses); err != nil {
		return
	}
	return
}

// AddMonitoredAddresses this endpoint takes new addresses and adds to monitor
//
// For more information: https://www.bitindex.network/developers/api-documentation-v3.html#Webhook
func (c *Client) AddMonitoredAddresses(addAddresses *MonitoredAddresses) (addresses MonitoredAddresses, err error) {

	var data []byte
	data, err = json.Marshal(addAddresses)
	if err != nil {
		return
	}

	var resp string
	// /api/v3/network/webhook/monitored_addrs
	resp, err = c.Request("webhook/monitored_addrs", http.MethodPut, data)
	if err != nil {
		return
	}

	addresses = *new(MonitoredAddresses)
	if err = json.Unmarshal([]byte(resp), &addresses); err != nil {
		return
	}
	return
}
