package states

import (
	cmds "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/commands"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
)

// State when bot waits for term
type TermState struct{}

func (s *TermState) PrepareMessage() string { return "Please enter a term you want to add" }

func (s *TermState) ProcessResponse(tc tm.TermitClient, u *Update) (*transition, error) {
	switch u.Cmd {
	case cmds.Reset:
		return NewTransition(&DefaultState{}), nil
	case cmds.Unknown:
	default:
		return NewTransition(s), TermExpectedError
	}
	return NewTransition(&TranslationState{term: u.Text}), nil
}
