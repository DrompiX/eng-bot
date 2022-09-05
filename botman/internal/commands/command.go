package states

import "fmt"

type Command string

const (
	Add     Command = "add"
	List    Command = "list"
	Task    Command = "task"
	Reset   Command = "reset"
	Unknown Command = "unknown"
)

var availableCommands = map[Command]struct{}{
	Add:   {},
	Task:  {},
	List:  {},
	Reset: {},
}

func ParseCommand(command string) (Command, error) {
	c := Command(command)
	if _, exists := availableCommands[c]; exists {
		return c, nil
	}
	return Unknown, fmt.Errorf("unknown command \"/%s\"", command)
}
