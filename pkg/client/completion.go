package client

import (
	"chatgpt-service/internal/pkg/constants"
)

const OpenAICompletionEndPoint = "/completions"
const CreateCompletionEndpoint = OpenAICompletionEndPoint + "/create"

// CompletionRequest is a request for the completions API
// https://github.com/PullRequestInc/go-gpt3/blob/283ab6b3e423c5567217fbe4e49950614ddd04c9/models.go#L36
type CompletionRequest struct {
	Model *string `json:"model"`
	// A list of string prompts to use.
	// TODO there are other prompt types here for using token integers that we could add support for.
	Prompt string `json:"prompt"`
	// How many tokens to complete up to. Max of 512
	MaxTokens *int `json:"max_tokens,omitempty"`
	// Sampling temperature to use
	Temperature *float32 `json:"temperature,omitempty"`
	// Whether to stream back results or not. Don't set this value in the request yourself
	// as it will be overriden depending on if you use CompletionStream or CreateCompletion methods.
	Stream *bool `json:"stream,omitempty"`
	// Alternative to temperature for nucleus sampling
	TopP *float32 `json:"top_p,omitempty"`
	// How many choice to create for each prompt
	N *int `json:"n"`
	// Include the probabilities of most likely tokens
	LogProbs *int `json:"logprobs"`
	// Echo back the prompt in addition to the completion
	Echo bool    `json:"echo"`
	User *string `json:"user"`
	// Up to 4 sequences where the API will stop generating tokens. Response will not contain the stop sequence.
	Stop []string `json:"stop,omitempty"`
	// PresencePenalty number between 0 and 1 that penalizes tokens that have already appeared in the text so far.
	PresencePenalty float32 `json:"presence_penalty"`
	// FrequencyPenalty number between 0 and 1 that penalizes tokens on existing frequency in the text so far.
	FrequencyPenalty float32 `json:"frequency_penalty"`
}

func NewCompletionRequest(prompt string, maxTokens int, model *string, stream *bool, temperature *float32) *CompletionRequest {
	cr := &CompletionRequest{
		Model:            new(string),
		Prompt:           prompt,
		MaxTokens:        new(int),
		Temperature:      new(float32),
		Stream:           new(bool),
		TopP:             new(float32),
		N:                new(int),
		LogProbs:         new(int),
		Echo:             false,
		Stop:             nil,
		PresencePenalty:  0.0,
		FrequencyPenalty: 0.0,
		User:             new(string),
	}
	if model != nil {
		cr.Model = model
	} else {
		*cr.Model = constants.TextDavinci003Engine
	}
	if stream != nil {
		cr.Stream = stream
	} else {
		*cr.Stream = false
	}
	if temperature != nil {
		cr.Temperature = temperature
	} else {
		*cr.Temperature = 0.0
	}
	*cr.MaxTokens = maxTokens
	*cr.TopP = 1.0
	*cr.N = 1
	*cr.LogProbs = 0
	*cr.User = constants.DefaultClientName
	return cr
}

// LogprobResult represents logprob result of Choice
type LogprobResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

type CompletionResponseChoice struct {
	Text         string        `json:"text"`
	Index        int           `json:"index"`
	LogProbs     LogprobResult `json:"logprobs"`
	FinishReason string        `json:"finish_reason"`
}

type CompletionResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// CompletionResponse is the full response from a request to the completions API
type CompletionResponse struct {
	ID      string                     `json:"id"`
	Object  string                     `json:"object"`
	Created int                        `json:"created"`
	Model   string                     `json:"model"`
	Choices []CompletionResponseChoice `json:"choices"`
	Usage   CompletionResponseUsage    `json:"usage"`
}
