package openai

import "time"

// ImageRequest is the request to send to the OpenAI API for image generation
// The only required field is the  prompt
// You should also set the user to the actual end-user that is sending the request
type ImageRequest map[string]any

// User is unique identifier representing your end-user, which will help OpenAI to monitor and detect abuse.
// You should set the user to the actual end-user that is sending the request
func (c ImageRequest) SetUser(user string) {
	c["user"] = user
}

func (c ImageRequest) SetPrompt(prompt string) {
	c["prompt"] = prompt
}

func (c ImageRequest) SetN(n string) {
	c["n"] = n
}

func (c ImageRequest) SetSize(size string) {
	c["size"] = size
}

func (c ImageRequest) SetFormat(format string) {
	c["response_format"] = format
}

type ImageResponse map[string]any

func (cr ImageResponse) URL() string {
	return cr["data"].([]any)[0].(map[string]any)["url"].(string)
}

func (cr ImageResponse) ImageBase64() string {
	return cr["data"].([]any)[0].(map[string]any)["b64_json"].(string)
}

func (cr ImageResponse) Created() time.Time {
	return time.Unix(int64(cr["created"].(float64)), 0)
}
