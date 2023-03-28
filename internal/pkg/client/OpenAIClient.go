package client

import (
	"bufio"
	"bytes"
	"chatgpt-service/pkg/client"
	cerror "chatgpt-service/pkg/errors"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const OpenAIClientKey = "OpenAIClient"
const GPTGenerateQueryEndpoint = "/gpt/generate"

type OpenAIClient struct {
	BaseURL       string
	ApiKey        string
	AccessToken   string
	UserAgent     string
	HttpClient    *http.Client
	DefaultEngine string
	IdOrg         string
}

func (oc *OpenAIClient) JSONBodyReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return bytes.NewBuffer(nil), nil
	}

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("failed to encode body: " + err.Error())
	}

	// Remove `OriginalPrompt` field if it exists
	var objMap map[string]interface{}
	if err := json.Unmarshal(raw, &objMap); err != nil {
		return nil, errors.New("failed to decode body: " + err.Error())
	}
	delete(objMap, "original_prompt")
	filteredRaw, err := json.Marshal(objMap)
	if err != nil {
		return nil, errors.New("failed to re-encode body: " + err.Error())
	}

	return bytes.NewBuffer(filteredRaw), nil
}

func (oc *OpenAIClient) NewRequestBuilder(ctx context.Context, method string, path string, payload interface{}) (*http.Request, error) {
	br, err := oc.JSONBodyReader(payload)
	if err != nil {
		return nil, err
	}
	url := oc.BaseURL + path // link to openai.com
	req, err := http.NewRequestWithContext(ctx, method, url, br)
	if err != nil {
		return nil, err
	}
	if len(oc.IdOrg) > 0 {
		req.Header.Set("user", oc.IdOrg)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oc.ApiKey))
	return req, nil
}

func (oc *OpenAIClient) ExecuteRequest(req *http.Request) (*http.Response, error) {
	resp, err := oc.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := oc.CheckRequestSucceed(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (oc *OpenAIClient) CheckRequestSucceed(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read from body: %w", err)
	}
	var result cerror.APIErrorResponse
	if err := json.Unmarshal(data, &result); err != nil {
		// if we can't decode the json error then create an unexpected error
		apiError := cerror.APIError{
			StatusCode: resp.StatusCode,
			Type:       "Unexpected",
			Message:    string(data),
		}
		return apiError
	}
	result.Error.StatusCode = resp.StatusCode
	return result.Error
}

func (oc *OpenAIClient) getResponseObject(rsp *http.Response, v interface{}) error {
	defer rsp.Body.Close()
	if err := json.NewDecoder(rsp.Body).Decode(v); err != nil {
		return fmt.Errorf("invalid json response: %w", err)
	}
	return nil
}

func (oc *OpenAIClient) ListModels(ctx context.Context) (*client.ListModelsResponse, error) {
	endPoint := client.ModelEndPoint
	req, err := oc.NewRequestBuilder(ctx, http.MethodGet, endPoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := oc.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(client.ListModelsResponse)
	if err := oc.getResponseObject(resp, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (oc OpenAIClient) RetrieveModel(ctx context.Context, engine string) (*client.ModelObject, error) {
	req, err := oc.NewRequestBuilder(ctx, http.MethodGet, client.ModelEndPoint+"/"+engine, nil)
	if err != nil {
		return nil, err
	}
	resp, err := oc.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(client.ModelObject)
	if err := oc.getResponseObject(resp, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (oc OpenAIClient) CreateNewChatCompletion(ctx context.Context, request client.ChatCompletionRequest) (*client.ChatCompletionResponse, error) {
	req, err := oc.NewRequestBuilder(ctx, http.MethodPost, client.OpenAIChatCompletionEndPoint, request)
	if err != nil {
		return nil, err
	}
	resp, err := oc.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(client.ChatCompletionResponse)
	if err := oc.getResponseObject(resp, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (oc OpenAIClient) CreateCompletion(ctx context.Context, request client.CompletionRequest) (*client.CompletionResponse, error) {
	return oc.CreateCompletionWithEngine(ctx, oc.DefaultEngine, request)
}

func (oc OpenAIClient) CompletionStream(ctx context.Context, request client.CompletionRequest, onData func(response *client.CompletionResponse)) error {
	return oc.CompletionStreamWithEngine(ctx, oc.DefaultEngine, request, onData)
}

func (oc OpenAIClient) CreateCompletionWithEngine(ctx context.Context, _ string, request client.CompletionRequest) (*client.CompletionResponse, error) {
	req, err := oc.NewRequestBuilder(ctx, http.MethodPost, client.OpenAICompletionEndPoint, request)
	if err != nil {
		return nil, err
	}
	resp, err := oc.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(client.CompletionResponse)
	if err := oc.getResponseObject(resp, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (oc OpenAIClient) CompletionStreamWithEngine(ctx context.Context, engine string, request client.CompletionRequest, onData func(response *client.CompletionResponse)) error {
	var dataPrefix = string([]byte("data: "))
	var doneSequence = string([]byte("[DONE]"))

	if !*request.Stream {
		return errors.New("stream option is false")
	}

	req, err := oc.NewRequestBuilder(ctx, http.MethodPost, client.OpenAICompletionEndPoint, request)
	if err != nil {
		return err
	}

	resp, err := oc.ExecuteRequest(req)
	if err != nil {
		return err
	}

	// Create a new channel to handle errors
	errCh := make(chan error)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("failed to close response body")
			// Send the error to the channel
			errCh <- err
		}
		// Close the channel
		close(errCh)
	}(resp.Body)

	// Create a new scanner to read the response body
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		// Moved the select statement inside the for loop to handle errors and context cancellation at every iteration.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return err
		default:
			line := scanner.Text()
			if !strings.HasPrefix(line, dataPrefix) {
				continue
			}
			line = strings.TrimPrefix(line, dataPrefix)
			if strings.HasPrefix(line, doneSequence) {
				break
			}
			output := new(client.CompletionResponse)
			if err := json.Unmarshal([]byte(line), output); err != nil {
				return errors.New("invalid json stream data: " + err.Error())
			}
			onData(output)
		}
	}

	if err := scanner.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to scan response body")
	}

	return nil
}

func (oc OpenAIClient) Edits(ctx context.Context, request client.EditsRequest) (*client.EditsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (oc OpenAIClient) Embeddings(ctx context.Context, request client.EmbeddingsRequest) (*client.EmbeddingsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (oc OpenAIClient) GetAccessToken() (string, error) {
	if oc.AccessToken == "" {
		return "", errors.New("access token is empty")
	}
	return oc.AccessToken, nil
}
