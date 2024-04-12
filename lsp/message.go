package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
	// params
	// it will be diff for diff Method
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
	// result
	// error
	// above 2 are diff to diff method..
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
