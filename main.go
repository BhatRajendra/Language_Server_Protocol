package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"lsp/analysis"
	"lsp/lsp"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("./log.txt")
	logger.Println("init")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	state := analysis.NewState()
	writer := os.Stdout
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.Decode_message(msg)
		if err != nil {
			logger.Printf("We got this error: %s\n", err)
			continue
		}
		handleMessage(logger, writer, state, method, content)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, content []byte) {
	logger.Printf("Received method: %s\n", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		// content is still in json format, only header was unmarshalled in rpc.Decode_message
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("could not unmarshal: %s\n", err)
		}
		logger.Printf("Connected to %s \t %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		// reply
		msg := lsp.NewinitializeResponse(request.ID)
		writeResponse(writer, msg)
		// didOpen
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		// content is still in json format, only header was unmarshalled in rpc.Decode_message
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("inside didOpen : could not unmarshal: %s\n", err)
		}
		logger.Printf("Opened :  to %s", request.Params.TextDocument.URI)
		// when a new doc is openend...we put it in analysis.state
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		// didChange
	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		// content is still in json format, only header was unmarshalled in rpc.Decode_message
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("inside didChange : could not unmarshal: %s\n", err)
		}
		logger.Printf("Changed :  to %s", request.Params.TextDocument.URI)
		// when a new doc is openend...we put it in analysis.state
		for _, change := range request.Params.ContentChange {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		// content is still in json format, only header was unmarshalled in rpc.Decode_message
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("inside hover : could not unmarshal: %s\n", err)
		}
		// create a response
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		// write it back
		writeResponse(writer, response)
	case "textDocument/definition":
		var request lsp.DefinitionRequest
		// content is still in json format, only header was unmarshalled in rpc.Decode_message
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("inside Defition : could not unmarshal: %s\n", err)
		}
		// create a response
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		// write it back
		writeResponse(writer, response)
	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		// content is still in json format, only header was unmarshalled in rpc.Decode_message
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("inside codeAction : could not unmarshal: %s\n", err)
		}
		// create a response
		response := state.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI)
		// write it back
		writeResponse(writer, response)
	case "textDocument/completion":
		var request lsp.CompletionRequest
		// content is still in json format, only header was unmarshalled in rpc.Decode_message
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("inside completion : could not unmarshal: %s\n", err)
		}
		// create a response
		response := state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)
		// write it back
		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.Encode_message(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me good file")
	}
	return log.New(logfile, "[LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
