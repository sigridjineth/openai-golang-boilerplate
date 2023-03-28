package client

import (
	"chatgpt-service/internal/pkg/engine"
	"fmt"
)

type Prompt struct {
	message string
	schema  engine.Schema
}

func (p *Prompt) String() string {
	message := p.message + "\n" + p.schema.String()
	return message
}

// TODO: only single prompt at the moment
func CreatePrompt(promptRaw GPTPromptRequest) (*Prompt, error) {
	prompt := promptRaw.Prompt
	message := "I want to you to act like the expert who are good at writing SQL query." + "\n"
	message += fmt.Sprintf("Given the table below, please write the SQL query that can get the info about %s", prompt) + "\n"
	message += "GIVE A SQL QUERY ONLY **without** any further responses/explanations." + "\n"
	message += "If you are not 100% certain to get the valid information from the database table below, respond \"NO\" without further responses/explanations.\n"
	message += "To query the database, you need to use the database information below"
	schema := engine.CreateEthereumCoreTransactionSchema()
	return &Prompt{message: message, schema: *schema}, nil
}
