package states

import (
	"context"
	"fmt"

	cmds "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/commands"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
)

// State when bot waits for an answer for a task
type AnswerState struct {
	tm.TermitTask
}

func (s *AnswerState) PrepareMessage() string {
	return fmt.Sprintf("Task to solve:\n%s -> ?", s.Term)
}

func (s *AnswerState) ProcessResponse(tc tm.TermitClient, u *Update) (*transition, error) {
	this := NewTransition(s)
	switch u.Cmd {
	case cmds.Reset:
		return NewTransition(&DefaultState{}), nil
	case cmds.Unknown:
	default:
		return this, TranslationExpectedError
	}

	ctx := tm.PrepareAuthCtx(context.Background(), u.ChatId)
	check, err := tc.CheckAnswer(ctx, tm.TermitTaskAnswer{ID: s.ID, Translation: u.Text})
	if err != nil {
		return this, err
	}

	res := NewTransition(&DefaultState{})
	if check.Success {
		res.Msg = "You are absolutely right. Good job!"
	} else {
		res.Msg = fmt.Sprintf(
			"Oops, answer is not correct\nReceived: %s, Expected: %s", check.Answer, check.Expect,
		)
	}
	return res, nil
}