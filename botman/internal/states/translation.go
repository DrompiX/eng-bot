package states

import (
	"context"

	cmds "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/commands"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
)

// State when bot waits for term's translation
type TranslationState struct {
	term string
}

func (s *TranslationState) PrepareMessage() string { return "Please enter a translation" }

func (s *TranslationState) ProcessResponse(tc tm.TermitClient, u *Update) (*transition, error) {
	this := NewTransition(s)
	switch u.Cmd {
	case cmds.Reset:
		return NewTransition(&DefaultState{}), nil
	case cmds.Unknown:
	default:
		return this, TranslationExpectedError
	}

	ctx := tm.PrepareAuthCtx(context.Background(), u.ChatId)
	err := tc.AddTerm(ctx, tm.TermitTerm{Term: s.term, Translation: u.Text})

	res := NewTransition(&DefaultState{})
	if err == nil {
		res.Msg = "Term was succesfully added to your collection!"
	}
	return res, err
}
