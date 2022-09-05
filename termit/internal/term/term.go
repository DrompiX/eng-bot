package term

import (
	"errors"
	"strings"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
)

type Term struct {
	Uid         user.UserID
	Data        string
	Translation string
}

var (
	ErrEmptyTerm        = errors.New("term can not be empty")
	ErrEmptyTranslation = errors.New("term has to contain non-empty translation")
)

func NewTerm(termData, translation string, uid user.UserID) (*Term, error) {
	if termData == "" {
		return nil, ErrEmptyTerm
	}
	if translation == "" {
		return nil, ErrEmptyTranslation
	}
	return &Term{uid, strings.ToLower(termData), strings.ToLower(translation)}, nil
}
