package lsp

// InitializeRequest
type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// now....there is lot lot more where that came from
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// initializeResponse
type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct{}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewinitializeResponse(id int) *InitializeResponse {
	return &InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{},
			ServerInfo: ServerInfo{
				Name:    "Go Language Server",
				Version: "v0.0.1",
			},
		},
	}
}
