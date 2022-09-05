package task

import (
	"github.com/google/uuid"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
)

type TaskID string

func NewTaskID() TaskID {
	return TaskID(uuid.NewString())
}

type task struct {
	ID       TaskID
	Uid      user.UserID
	Term     string
	Expected string
	Success  *bool
}

func NewTask(uid user.UserID, term string, translation string) *task {
	return &task{ID: NewTaskID(), Uid: uid, Term: term, Expected: translation, Success: nil}
}

type answer struct {
	Tid         TaskID
	Uid         user.UserID
	Translation string
}

func NewAnswer(tid TaskID, uid user.UserID, translation string) *answer {
	return &answer{Tid: tid, Uid: uid, Translation: translation}
}

type AnswerCheck struct {
	Success  bool
	Expected string
}
