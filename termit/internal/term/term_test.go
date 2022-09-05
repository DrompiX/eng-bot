package term

import (
	"testing"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/utils"
)

func TestNewTerm(t *testing.T) {
	type expectedResult struct {
		*Term
		error
	}

	tests := []struct {
		name        string
		termData    string
		translation string
		expected    expectedResult
	}{
		{
			name:     "Empty term can not be created",
			expected: expectedResult{nil, ErrEmptyTerm},
		},
		{
			name:        "Term can not have empty translation",
			termData:    "hello",
			translation: "",
			expected:    expectedResult{nil, ErrEmptyTranslation},
		},
		{
			name:        "Successfully create term",
			termData:    "hello",
			translation: "привет",
			expected:    expectedResult{
				&Term{Data: "hello", Translation: "привет", Uid: "testuser"},
				nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			term, err := NewTerm(test.termData, test.translation, "testuser")
			utils.Equals(t, test.expected.error, err)
			utils.Equals(t, test.expected.Term, term)
		})
	}
}
