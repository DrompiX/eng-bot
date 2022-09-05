package states

import (
	"context"
	"fmt"
	"strings"

	cmds "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/commands"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
)

// Initial state of bot, accepts any command
type DefaultState struct{}

func (s *DefaultState) PrepareMessage() string { return "Ready to accept commands" }

func (s *DefaultState) ProcessResponse(tc tm.TermitClient, u *Update) (*transition, error) {
	this := NewTransition(s)
	switch u.Cmd {
	case cmds.Add:
		return NewTransition(&TermState{}), nil

	case cmds.List:
		ctx := tm.PrepareAuthCtx(context.Background(), u.ChatId)
		terms, err := tc.GetCollection(ctx)
		if err != nil {
			return this, err
		}

		// Generate collection message
		sb := strings.Builder{}
		sb.WriteString("Collection of your terms:\n\n")
		for i, t := range terms {
			sb.WriteString(fmt.Sprintf("%d) %s -> %s\n\n", i, t.Term, t.Translation))
		}
		this.Msg = sb.String()

		return this, nil

	case cmds.Task:
		ctx := tm.PrepareAuthCtx(context.Background(), u.ChatId)
		task, err := tc.GetTask(ctx)
		if err != nil {
			return this, err
		}
		return NewTransition(&AnswerState{task}), nil

	case cmds.Reset:
		return this, AlreadyAtStartError

	case cmds.Unknown:
	}
	return this, CommandExpectedError
}
