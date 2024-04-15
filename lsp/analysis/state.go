package analysis

type State struct {
	// Documents is map which maps openend document to its content
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: map[string]string{},
	}
}

func (s *State) OpenDocument(document, text string) {
	s.Documents[document] = text
}
