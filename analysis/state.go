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

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(text, "VS Code") {
			idx := strings.Index(line, "VS Code")
			if idx > -1 {
				diagnostics = append(diagnostics, lsp.Diagnostic{
					Range:    LineRange(row, idx, idx+len("VS Code")),
					Severity: 1,
					Source:   "Common Sense",
					Message:  "Please make use of some good editors like nvim to make your life easier",
				})
			}
		}
		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line, "Neovim")
			if idx > -1 {
				diagnostics = append(diagnostics, lsp.Diagnostic{
					Range:    LineRange(row, idx, idx+len("Neovim")),
					Severity: 3,
					Source:   "Common Sense",
					Message:  "Great choice :)",
				})
			}
		}
	}
	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
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

func (s *State) TextDocumentCompletion(id int, URI string) lsp.CompletionResponse {
	// ask you static analysis tools to figure out the best TextDocumentCompletion
	// here the response is very very very minimal
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim BTW",
			Detail:        "Very cool editor",
			Documentation: "It's fun to write in Nvim and all credit goes to TJ DeVries",
		},
		{
			Label:         "I am Rajendra",
			Detail:        "Amatuer programmer",
			Documentation: "but stillllll....it's fun to programme stuff",
		},
	}
	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{Line: line, Character: start},
		End:   lsp.Position{Line: line, Character: end},
	}
}
