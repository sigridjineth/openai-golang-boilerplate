package client

type GPTPromptRequest struct {
	Prompt string `json:"prompt"`
}

type GPTPromptSuccessfulResponse struct {
	Result string `json:"result"`
}
