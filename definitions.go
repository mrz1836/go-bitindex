package bitindex

// NetworkType is used internally to represent the possible values
// for network in queries to be submitted: {"main", "test", "stn"}
type NetworkType string

// APIInternalError is for internal server errors
type APIInternalError struct {
	Errors       []string `json:"errors,omitempty"`
	ErrorMessage string   `json:"message,omitempty"`
	ErrorName    string   `json:"name,omitempty"`
}

// APIErrorResponse is from bitindex (broadcast related errors)
type APIErrorResponse struct {
	Errors  []string        `json:"errors,omitempty"`
	Message APIErrorMessage `json:"message,omitempty"`
}

// APIErrorMessage is the nested error from APIErrorResponse
type APIErrorMessage struct {
	ErrorCode    int    `json:"code,omitempty"`
	ErrorMessage string `json:"message,omitempty"`
}

// AddressInfo is the address info for a returned address request
type AddressInfo struct {
	APIInternalError
	Address                    string   `json:"addrStr"`
	Balance                    float64  `json:"balance"`
	BalanceSatoshis            int64    `json:"balanceSat"`
	TotalReceived              float64  `json:"totalReceived"`
	TotalReceivedSatoshis      int64    `json:"totalReceivedSat"`
	TotalSent                  float64  `json:"totalSent"`
	TotalSentSatoshis          int64    `json:"totalSentSat"`
	Transactions               []string `json:"transactions"`
	TxAppearances              int64    `json:"txApperances"`
	UnconfirmedBalance         float64  `json:"unconfirmedBalance"`
	UnconfirmedBalanceSatoshis int64    `json:"unconfirmedBalanceSat"`
	UnconfirmedTxAppearances   int64    `json:"unconfirmedTxApperances"`
}

// UnspentTransactions is a list of unspent transactions
type UnspentTransactions []UnspentTransaction

// UnspentTransaction is a standard UTXO response
type UnspentTransaction struct {
	Address       string  `json:"address"`
	Amount        float64 `json:"amount"`
	Confirmations int64   `json:"confirmations"`
	Height        int64   `json:"height"`
	OutputIndex   int64   `json:"outputIndex"`
	Satoshis      int64   `json:"satoshis"`
	Script        string  `json:"script"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	TxID          string  `json:"txid"`
	Value         int64   `json:"value"`
	Vout          int     `json:"vout"`
}

// GetTransactionsRequest is for making a POST to get transactions
type GetTransactionsRequest struct {
	Address        string   `json:"addrs"` // single address or addr1,addr2,addr3
	Addresses      []string `json:"-"`     // (used for multiple)
	AfterBlockHash string   `json:"afterBlockHash,omitempty"`
	AfterHeight    string   `json:"afterHeight,omitempty"`
	FromIndex      int64    `json:"fromIndex,omitempty"`
	IncludeAsm     bool     `json:"includeAsm"`
	IncludeHex     bool     `json:"includeHex"`
	ToIndex        int64    `json:"toIndex,omitempty"`
}

// GetTransactionsResponse is the response from the POST request
type GetTransactionsResponse struct {
	APIInternalError
	TotalItems int64         `json:"totalItems"`
	From       int64         `json:"from"`
	To         int64         `json:"to"`
	Items      []Transaction `json:"items"`
}

// Transaction is returned in the GetTransactionsResponse
type Transaction struct {
	APIInternalError
	BlockHash     string       `json:"blockhash"`
	BlockHeight   int64        `json:"blockheight"`
	BlockTime     int64        `json:"blocktime"`
	Confirmations int64        `json:"confirmations"`
	Fees          float64      `json:"fees"`
	Hash          string       `json:"hash"`
	LockTime      int64        `json:"locktime"`
	RawTx         string       `json:"rawtx"`
	Size          int64        `json:"size"`
	Time          int64        `json:"time"`
	TxID          string       `json:"txid"`
	ValueIn       float64      `json:"valueIn"`
	ValueOut      float64      `json:"valueOut"`
	Version       int          `json:"version"`
	Vin           []VinObject  `json:"vin"`
	Vout          []VoutObject `json:"vout"`
}

// TransactionRaw is the response for the raw tx request
type TransactionRaw struct {
	APIInternalError
	RawTx string `json:"rawtx"`
}

// VinObject is the vin data
type VinObject struct {
	Address       string          `json:"address"`
	AddressAddr   string          `json:"addr"`
	N             int             `json:"n"`
	ScriptSig     ScriptSigObject `json:"scriptSig"`
	Sequence      int64           `json:"sequence"`
	TxID          string          `json:"txid"`
	Value         float64         `json:"value"`
	ValueSatoshis int64           `json:"valueSat"`
	Vout          int             `json:"vout"`
}

// ScriptSigObject is the script signature data
type ScriptSigObject struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

// VoutObject is the vout data
type VoutObject struct {
	N             int                `json:"n"`
	ScriptPubKey  ScriptPubKeyObject `json:"scriptPubKey"`
	SpentHeight   int64              `json:"spentHeight"`
	SpentIndex    int64              `json:"spentIndex"`
	SpentTxID     string             `json:"spentTxId"`
	Value         float64            `json:"value"`
	ValueSatoshis int64              `json:"valueSat"`
}

// ScriptPubKeyObject is the script pubkey data
type ScriptPubKeyObject struct {
	Addresses          []string `json:"addresses"`
	Asm                string   `json:"asm"`
	Hex                string   `json:"hex"`
	RequiredSignatures int      `json:"reqSigs"`
	Type               string   `json:"type"`
}

// GetUnspentTransactionsRequest is for making the GetUnspentTransactions request
type GetUnspentTransactionsRequest struct {
	Address   string   `json:"addrs"` // single address or addr1,addr2,addr3
	Addresses []string `json:"-"`     // (used for multiple)
	Sort      string   `json:"sort"`  // Format is 'field:asc' such as 'value:desc' to sort by value descending
}

// SendTransactionResponse is the response for the request
type SendTransactionResponse struct {
	APIErrorResponse
	TxID string `json:"txid"`
}
