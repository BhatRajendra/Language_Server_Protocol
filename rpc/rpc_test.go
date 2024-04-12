package rpc_test

import (
	"lsp/rpc"
	"testing"
)

type EncodeExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.Encode_message(EncodeExample{Testing: true})
	if actual != expected {
		t.Fatalf("Expected %s, but got %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.Decode_message([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected Content-Length to be 15, but got %d", contentLength)
	}
	if method != "hi" {
		t.Fatalf("Expected method to be hi, but got %s", method)
	}
}
