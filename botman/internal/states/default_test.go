package states

import (
	"testing"

	cmds "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/commands"
	tm "gitlab.ozon.dev/DrompiX/homework-2/botman/internal/termit"
	"gitlab.ozon.dev/DrompiX/homework-2/botman/internal/utils"
)

type Updates []*Update

type stateTest struct {
	name        string
	actions     Updates
	expectedRes *transition
	expectedErr error
}

func TestDefaultState(t *testing.T) {
	var defaultStateTests = []stateTest{
		{
			name: "Expects command in default state",
			actions: Updates{
				&Update{Cmd: cmds.Unknown, Text: ""},
				&Update{Cmd: cmds.Unknown, Text: "some random /text"},
			},
			expectedRes: &transition{New: &DefaultState{}},
			expectedErr: CommandExpectedError,
		},
		{
			name:        "Tranitions to TermState on Add command",
			actions:     Updates{&Update{Cmd: cmds.Add}},
			expectedRes: &transition{New: &TermState{}},
			expectedErr: nil,
		},
		// TODO: generate mock for TermitClient and refactor tests
		// {
		// 	name:        "Returns collection of terms on List command",
		// 	actions:     Updates{&update{Cmd: List}},
		// 	expectedRes: &transition{New: &DefaultState{}, Msg: "Collection of your terms:\n\n"},
		// 	expectedErr: nil,
		// },
		// {
		// 	name:        "Returns task on Task command",
		// 	actions:     Updates{&update{Cmd: Task}},
		// 	expectedRes: &transition{New: &DefaultState{}, Msg: "Task to solve: hello -> ?"},
		// 	expectedErr: nil,
		// },
		{
			name:        "Reset does not perform any transition",
			actions:     Updates{&Update{Cmd: cmds.Reset}},
			expectedRes: &transition{New: &DefaultState{}},
			expectedErr: AlreadyAtStartError,
		},
	}

	termitClient := &tm.GrpcTermitClient{
		TermServ: nil,
		TaskServ: nil,
	}

	for _, test := range defaultStateTests {
		t.Run(test.name, func(t *testing.T) {
			for _, action := range test.actions {
				startState := &DefaultState{}
				res, err := startState.ProcessResponse(termitClient, action)
				utils.Equals(t, test.expectedErr, err)
				utils.Equals(t, test.expectedRes, res)
			}
		})
	}
}
