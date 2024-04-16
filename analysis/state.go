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

func (s *State) Definition(id int, URI string, pos lsp.Position) lsp.DefinitionResponse {
	// in real life it worlds would be more complex
	// in real life it would show or go to the Definition
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: URI,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      pos.Line - 2,
					Character: 0,
				},
				End: lsp.Position{
					Line:      pos.Line - 2,
					Character: 0,
				},
			},
		},
	}
}
