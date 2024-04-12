package main

import (
	"bufio"
	"encoding/json"
	"log"
	"lsp/lsp"
	"lsp/rpc"
	"os"
)

// TODO: uncomment lsp server in nvim config while working on this

func main() {
	logger := getLogger("./log.txt")
	logger.Println("init")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.Decode_message(msg)
		if err != nil {
			logger.Printf("We got this error: %s\n", err)
			continue
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
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
		reply := rpc.Encode_message(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Println("reply sent")
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me good file")
	}
	return log.New(logfile, "[LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
