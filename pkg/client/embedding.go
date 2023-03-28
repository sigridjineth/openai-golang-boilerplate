package client

// EmbeddingsRequest is a request for the Embeddings API
type EmbeddingsRequest struct {
	// Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings
	// for multiple inputs in a single request, pass an array of strings or array of token arrays.
	// Each input must not exceed 2048 tokens in length.
	Input []string `json:"input"`
	// ID of the model to use
	Model string `json:"model"`
	// The request user is an optional parameter meant to be used to trace abusive requests
	// back to the originating user. OpenAI states:
	// "The [user] IDs should be a string that uniquely identifies each user. We recommend hashing
	// their username or email address, in order to avoid sending us any identifying information.
	// If you offer a preview of your product to non-logged in users, you can send a session ID
	// instead."
	User string `json:"user,omitempty"`
}

// The inner result of a create embeddings request, containing the embeddings for a single input.
type EmbeddingsResult struct {
	// The type of object returned (e.g., "list", "object")
	Object string `json:"object"`
	// The embedding data for the input
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// The usage stats for an embeddings response
type EmbeddingsUsage struct {
	// The number of tokens used by the prompt
	PromptTokens int `json:"prompt_tokens"`
	// The total tokens used
	TotalTokens int `json:"total_tokens"`
}

// EmbeddingsResponse is the response from a create embeddings request.
// https://beta.openai.com/docs/api-reference/embeddings/create
type EmbeddingsResponse struct {
	Object string             `json:"object"`
	Data   []EmbeddingsResult `json:"data"`
	Usage  EmbeddingsUsage    `json:"usage"`
}
