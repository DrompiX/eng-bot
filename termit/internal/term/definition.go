package term

import "errors"

var (
	ErrEmptyDefinition = errors.New("definition can not be empty")
)

type Definition struct {
	Info    string
	Example *string
}

func NewDefinition(info string, example *string) (*Definition, error) {
	if info == "" {
		return nil, ErrEmptyDefinition
	}
	return &Definition{info, example}, nil
}
