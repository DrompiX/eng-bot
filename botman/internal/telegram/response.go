package telegram

type botResponse struct {
	Messages []string
}

func newBotResponse() *botResponse {
	return &botResponse{Messages: []string{}}
}

func (r *botResponse) addMessage(msg string) {
	if msg != "" {
		r.Messages = append(r.Messages, msg)
	}
}