package openai

import "time"

// CompletionRequest is the request to send to the OpenAI API for completion
// The 2 required fields are prompt and model
// You should also set the user to the actual end-user that is sending the request
type CompletionRequest map[string]any

// User is unique identifier representing your end-user, which will help OpenAI to monitor and detect abuse.
// You should set the user to the actual end-user that is sending the request
// OpenAI has since changed their policy to not require a user
func (c CompletionRequest) SetUser(user string) {
	c["user"] = user
}

func (c CompletionRequest) SetModel(model string) {
	c["model"] = model
}

func (c CompletionRequest) SetPrompt(prompt string) {
	c["prompt"] = prompt
}

type CompletionResponse map[string]any

func (cr CompletionResponse) Text() string {
	return cr["choices"].([]any)[0].(map[string]any)["text"].(string)
}

func (cr CompletionResponse) Id() string {
	return cr["id"].(string)
}

func (cr CompletionResponse) Model() string {
	return cr["model"].(string)
}

func (cr CompletionResponse) Object() string {
	return cr["object"].(string)
}

func (cr CompletionResponse) Created() time.Time {
	return time.Unix(int64(cr["created"].(float64)), 0)
}

func (cr CompletionResponse) Choices() map[string]any {
	return cr["choices"].([]any)[0].(map[string]any)
}

func (cr CompletionResponse) Usage() map[string]any {
	return cr["usage"].(map[string]any)
}
