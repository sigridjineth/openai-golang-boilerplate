package client

import (
	"chatgpt-service/pkg/client"
	"context"
	"github.com/labstack/echo/v4"
)

type EchoHttpMethodFunc func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

// OpenAIClientInterface is an API client to communicate with the OpenAI gpt-3 APIs
// https://github.com/PullRequestInc/go-gpt3/blob/283ab6b3e423c5567217fbe4e49950614ddd04c9/gpt3.go
type OpenAIClientInterface interface {
	// ListModels Lists the currently available models
	// and provides basic information about each one such as the owner and availability.
	// curl https://api.openai.com/v1/models -H 'Authorization: Bearer YOUR_API_KEY'
	ListModels(ctx context.Context) (*client.ListModelsResponse, error)

	// RetrieveModel retrieves a client instance, providing basic information about the client such
	// as the owner and availability.
	RetrieveModel(ctx context.Context, engine string) (*client.ModelObject, error)

	// Completion creates a completion with the default client. This is the main endpoint of the API
	// which auto-completes based on the given prompt.
	CreateCompletion(ctx context.Context, request client.CompletionRequest) (*client.CompletionResponse, error)

	// CompletionStream creates a completion with the default client and streams the results through
	// multiple calls to onData.
	CompletionStream(ctx context.Context, request client.CompletionRequest, onData func(response *client.CompletionResponse)) error

	// CompletionWithEngine is the same as Completion except allows overriding the default client on the client
	CreateCompletionWithEngine(ctx context.Context, engine string, request client.CompletionRequest) (*client.CompletionResponse, error)

	// CompletionStreamWithEngine is the same as CompletionStream except allows overriding the default client on the client
	CompletionStreamWithEngine(ctx context.Context, engine string, request client.CompletionRequest, onData func(response *client.CompletionResponse)) error

	// Given a prompt and an instruction, the client will return an edited version of the prompt.
	Edits(ctx context.Context, request client.EditsRequest) (*client.EditsResponse, error)

	// Returns an embedding using the provided request.
	Embeddings(ctx context.Context, request client.EmbeddingsRequest) (*client.EmbeddingsResponse, error)
}
