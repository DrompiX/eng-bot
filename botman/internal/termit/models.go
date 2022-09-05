package termit

type TermitTerm struct {
	Term       string
	Translation string
}

type TermitTask struct {
	ID   string
	Term string
}

type TermitTaskAnswer struct {
	ID          string
	Translation string
}

type ValidatedAnswer struct {
	Success bool
	Answer  string
	Expect  string
}
