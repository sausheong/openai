package openai

import "time"

type CompletionRequest map[string]any

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
