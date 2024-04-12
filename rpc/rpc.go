package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func Encode_message(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func Decode_message(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("did not find the separator")
	}
	contentLengthByte := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthByte))
	if err != nil {
		return "", nil, err
	}
	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}
	return baseMessage.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	contentLengthByte := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthByte))
	if err != nil {
		return 0, nil, err
	}
	if len(content) < contentLength {
		return 0, nil, nil
	}
	totalLength := len(header) + 4 + contentLength // 4 is for '\r\n\r\n'
	return totalLength, data[0:totalLength], nil
}
