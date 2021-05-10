package bitindex

// NetworkType is used internally to represent the possible values
// for network in queries to be submitted: {"main", "test", "stn"}
type NetworkType string

const (
	// apiKeyField is the field for the api key in all requests
	apiKeyField string = "api_key"

	// NetworkMain is for main-net
	NetworkMain NetworkType = "main"

	// NetworkTest is for test-net
	NetworkTest NetworkType = "test"

	// NetworkStn is for the stn-net
	NetworkStn NetworkType = "stn"
)

// APIInternalError is for internal server errors (most requests)
type APIInternalError struct {
	Errors       []string `json:"errors,omitempty"`
	ErrorMessage string   `json:"message,omitempty"`
	ErrorName    string   `json:"name,omitempty"`
}

// APIErrorResponse is from bitindex (broadcast related errors)
type APIErrorResponse struct {
	Error        string   `json:"error,omitempty"`
	ErrorCode    int      `json:"code,omitempty"`
	ErrorMessage string   `json:"message,omitempty"`
	Errors       []string `json:"errors,omitempty"`
	Success      bool     `json:"success,omitempty"`
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
// Also has some fields for xpub data (chain, num, path)
type UnspentTransaction struct {
	Address       string  `json:"address"`
	Amount        float64 `json:"amount"`
	Chain         int     `json:"chain"`
	Confirmations int64   `json:"confirmations"`
	Height        int64   `json:"height"`
	Num           int     `json:"num"`
	OutputIndex   int64   `json:"outputIndex"`
	Path          string  `json:"path"`
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
	Vin           []vinObject  `json:"vin"`
	Vout          []voutObject `json:"vout"`
}

// TransactionRaw is the response for the raw tx request
type TransactionRaw struct {
	APIInternalError
	RawTx string `json:"rawtx"`
}

// vinObject is the vin data
type vinObject struct {
	Address       string          `json:"address"`
	AddressAddr   string          `json:"addr"`
	N             int             `json:"n"`
	ScriptSig     scriptSigObject `json:"scriptSig"`
	Sequence      int64           `json:"sequence"`
	TxID          string          `json:"txid"`
	Value         float64         `json:"value"`
	ValueSatoshis int64           `json:"valueSat"`
	Vout          int             `json:"vout"`
}

// scriptSigObject is the script signature data
type scriptSigObject struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

// voutObject is the vout data
type voutObject struct {
	N             int                `json:"n"`
	ScriptPubKey  scriptPubKeyObject `json:"scriptPubKey"`
	SpentHeight   int64              `json:"spentHeight"`
	SpentIndex    int64              `json:"spentIndex"`
	SpentTxID     string             `json:"spentTxId"`
	Value         float64            `json:"value"`
	ValueSatoshis int64              `json:"valueSat"`
}

// scriptPubKeyObject is the script pubkey data
type scriptPubKeyObject struct {
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

// ChainInfoResponse response struct for chain info request
type ChainInfoResponse struct {
	Info chainInfo `json:"info"`
}

// chainInfo is for the chain info request response data
type chainInfo struct {
	Blocks          int64   `json:"blocks"`
	Connections     int64   `json:"connections"`
	Difficulty      float64 `json:"difficulty"`
	Errors          string  `json:"errors"`
	Network         string  `json:"network"`
	ProtocolVersion int64   `json:"protocolversion"`
	Proxy           string  `json:"proxy"`
	RelayFee        float64 `json:"relayfee"`
	TestNet         bool    `json:"testnet"`
	TimeOffset      int64   `json:"timeoffset"`
	Version         int64   `json:"version"`
}

// ChainDifficultyResponse response struct for chain difficulty request
type ChainDifficultyResponse struct {
	Difficulty float64 `json:"difficulty"`
}

// ChainBestBlockHashResponse response struct for best block hash request
type ChainBestBlockHashResponse struct {
	BestBlockHash string `json:"bestblockhash"`
}

// ChainLastBlockHashResponse response struct for last block hash request
type ChainLastBlockHashResponse struct {
	LastBlockHash string `json:"lastblockhash"`
	SyncTipHash   string `json:"syncTipHash"`
}

// BlockHashByHeightResponse response struct for block hash by height request
type BlockHashByHeightResponse struct {
	APIInternalError
	BlockHash string `json:"blockHash"`
}

// BlockHeaderResponse is the block header response
type BlockHeaderResponse struct {
	APIInternalError
	Bits              string  `json:"bits"`
	ChainWork         string  `json:"chainwork"`
	Confirmations     int64   `json:"confirmations"`
	Difficulty        float64 `json:"difficulty"`
	Hash              string  `json:"hash"`
	Height            int64   `json:"height"`
	MedianTime        int64   `json:"mediantime"`
	MerkleRoot        string  `json:"merkleroot"`
	NextBlockHash     string  `json:"nextblockhash"`
	Nonce             int64   `json:"nonce"`
	PreviousBlockHash string  `json:"previousblockhash"`
	Time              int64   `json:"time"`
	Version           int     `json:"version"`
	VersionHex        string  `json:"versionHex"`
}

// BlockResponse is the block response
type BlockResponse struct {
	APIInternalError
	Bits              string   `json:"bits"`
	ChainWork         string   `json:"chainwork"`
	Confirmations     int64    `json:"confirmations"`
	Difficulty        float64  `json:"difficulty"`
	Hash              string   `json:"hash"`
	Height            int64    `json:"height"`
	MedianTime        int64    `json:"mediantime"`
	MerkleRoot        string   `json:"merkleroot"`
	NextBlockHash     string   `json:"nextblockhash"`
	Nonce             int64    `json:"nonce"`
	PreviousBlockHash string   `json:"previousblockhash"`
	Size              int64    `json:"size"`
	Time              int64    `json:"time"`
	Tx                []string `json:"tx"`
	Version           int      `json:"version"`
	VersionHex        string   `json:"versionHex"`
}

// BlockRawResponse response struct for raw block request
type BlockRawResponse struct {
	APIInternalError
	RawBlock string `json:"rawblock"`
}

// WebhookConfigResponse is the response from get config
type WebhookConfigResponse struct {
	APIInternalError
	Enabled bool   `json:"enabled"`
	ID      string `json:"id"`
	Secret  string `json:"secret"`
	URL     string `json:"url"`
}

// WebhookUpdateConfig is for updating a webhook config
type WebhookUpdateConfig struct {
	Enabled bool   `json:"enabled"`
	Secret  string `json:"secret,omitempty"`
	URL     string `json:"url,omitempty"`
}

// MonitoredAddresses is the response from get monitored addresses
type MonitoredAddresses []MonitoredAddress

// MonitoredAddress is the address from get monitored addresses
type MonitoredAddress struct {
	Address string `json:"addr"`
}

// XpubAddresses is the list of next addresses
type XpubAddresses []XPubAddress

// XPubAddress is an address returned from xpub requests
type XPubAddress struct {
	Address string `json:"address"`
	Chain   int    `json:"chain"`
	Height  int64  `json:"height"`
	Num     int    `json:"num"`
	Path    string `json:"path"`
	TxID    string `json:"txid"`
}

// XpubBalance is the balance returned for a xpub address
type XpubBalance struct {
	APIInternalError
	Confirmed   int64 `json:"confirmed"`
	UnConfirmed int64 `json:"unconfirmed"`
}
