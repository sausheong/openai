package openai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewClient(auth, org string) *Client {
	return &Client{
		Authorization: auth,
		Organization:  org,
	}
}

func (c *Client) Complete(request CompletionRequest) (response CompletionResponse, err error) {
	// unmarshal the request to send to the API
	body, err := json.Marshal(request)
	if err != nil {
		return
	}
	// create a HTTP POST request
	req, err := http.NewRequest("POST", COMPLETIONS_URL, bytes.NewReader(body))
	req.Header.Add("Authorization", "Bearer "+c.Authorization)
	req.Header.Add("OpenAI-Organization", c.Organization)
	req.Header.Add("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	// create a HTTP client and use it to send the request
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// read body from HTTP response and store into the client
	body, err = ioutil.ReadAll(resp.Body)
	response = make(CompletionResponse)
	err = json.Unmarshal(body, &response)
	return
}
