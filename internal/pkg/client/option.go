package client

// ClientOption are options that can be passed when creating a new client
type ClientOption func(*OpenAIClient) error

// WithOrg is a client option that allows you to override the organization ID
func WithOrg(id string) ClientOption {
	return func(c *OpenAIClient) error {
		c.IdOrg = id
		return nil
	}
}

// WithDefaultEngine is a client option that allows you to override the default client of the client
func WithDefaultEngine(engine string) ClientOption {
	return func(c *OpenAIClient) error {
		c.DefaultEngine = engine
		return nil
	}
}

// WithUserAgent is a client option that allows you to override the default user agent of the client
func WithUserAgent(userAgent string) ClientOption {
	return func(c *OpenAIClient) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithBaseURL is a client option that allows you to override the default base url of the client.
// The default base url is "https://api.openai.com/v1"
func WithBaseURL(baseURL string) ClientOption {
	return func(c *OpenAIClient) error {
		c.BaseURL = baseURL
		return nil
	}
}
