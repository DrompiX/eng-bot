package states

import (
	"errors"

	cmds "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/commands"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
)

var (
	CommandExpectedError     = errors.New("enter a command, use /help to find available commands")
	TermExpectedError        = errors.New("specify a term or call /reset to reset state")
	TranslationExpectedError = errors.New("specify a translation or call /reset to reset state")
	AlreadyAtStartError      = errors.New("you are already in default state")
)

type State interface {
	PrepareMessage() string
	ProcessResponse(tm.TermitClient, *Update) (*transition, error)
}

// Internal representation of bot update command
type Update struct {
	Cmd    cmds.Command
	Text   string
	ChatId int64
}

func NewUpdate(chatId int64) *Update {
	return &Update{Cmd: cmds.Unknown, Text: "", ChatId: chatId}
}

// Represents transtion to New state with message as side effect
type transition struct {
	New State
	Msg string
}

func NewTransition(newState State) *transition {
	return &transition{New: newState, Msg: ""}
}
