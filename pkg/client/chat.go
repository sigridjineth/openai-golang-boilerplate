package client

import (
	"chatgpt-service/internal/pkg/constants"
	"chatgpt-service/internal/pkg/process"
)

const OpenAIChatCompletionEndPoint = "/chat/completions"

type ChatCompletionRequest struct {
	Model            *string     `json:"model"`
	Messages         []Message   `json:"messages"`
	MaxTokens        *int        `json:"max_tokens"`
	Temperature      *int        `json:"temperature"`
	Stream           *bool       `json:"stream"`
	TopP             *int        `json:"top_p"`
	N                *int        `json:"n"`
	Stop             interface{} `json:"stop"`
	PresencePenalty  float32     `json:"presence_penalty"`
	FrequencyPenalty float32     `json:"frequency_penalty"`
	User             *string     `json:"user"`
	OriginalPrompt   *string     `json:"original_prompt"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewChatCompletionRequest(content string, maxTokens int, model *string, stream *bool, role *string) (*ChatCompletionRequest, error) {
	promptRaw := GPTPromptRequest{
		Prompt: content,
	}
	prompt, err := CreatePrompt(promptRaw)
	if err != nil {
		return nil, err
	}
	promptInString := prompt.String()

	if role == nil {
		role = new(string)
		*role = "user"
	}
	cr := &ChatCompletionRequest{
		Model: new(string),
		Messages: []Message{
			{
				Role:    *role,
				Content: promptInString,
			},
		},
		MaxTokens:        new(int),
		Temperature:      new(int),
		Stream:           new(bool),
		TopP:             new(int),
		N:                new(int),
		Stop:             nil,
		PresencePenalty:  0.0,
		FrequencyPenalty: 0.0,
		User:             new(string),
		OriginalPrompt:   new(string),
	}
	if model != nil {
		cr.Model = model
	} else {
		*cr.Model = constants.Gpt35EngineTurbo
	}
	if stream != nil {
		cr.Stream = stream
	} else {
		*cr.Stream = false
	}
	*cr.MaxTokens = maxTokens
	*cr.Temperature = 0.0
	*cr.TopP = 1.0
	*cr.N = 1
	*cr.User = constants.DefaultClientName
	*cr.OriginalPrompt = content

	return cr, nil
}

type ChatCompletionResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	FinishReason string `json:"finish_reason"`
	Index        int    `json:"index"`
}

// GetContent TODO: temporarily only return the first choice
func (c *ChatCompletionRequest) GetContent() string {
	return process.ConvertNewLineToSpace(c.Messages[0].Content)
}

// GetContent TODO: temporarily only return the first choice
func (c *ChatCompletionResponse) GetContent() string {
	return process.ConvertNewLineToSpace(c.Choices[0].Message.Content)
}
