package analysis

import (
	"fmt"
	"lsp/lsp"
	"strings"
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

func (s *State) TextDocumentCodeAction(id int, URI string) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[URI]
	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[URI] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "I can change this bit of text",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Override VS Code",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[URI] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C**E",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Censor VS Code",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}
	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{Line: line, Character: start},
		End:   lsp.Position{Line: line, Character: end},
	}
}
