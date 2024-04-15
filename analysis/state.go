package analysis

import (
	"fmt"
	"lsp/lsp"
)

type State struct {
	// Documents is map which maps openend document to its content
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: map[string]string{},
	}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, URI string, pos lsp.Position) lsp.HoverResponse {
	doc := s.Documents[URI]
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File %s, Character %d", URI, len(doc)),
		},
	}
}
