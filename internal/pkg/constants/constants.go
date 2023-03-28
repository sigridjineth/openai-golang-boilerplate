package constants

import "fmt"

// GPT 3.5 model
const (
	Gpt35EngineTurbo     = "gpt-3.5-turbo"
	Gpt35EngineTurbo0301 = "gpt-3.5-turbo-0301"
	TextDavinci003Engine = "text-davinci-003"
)

// Engine Types
const (
	TextAda001Engine     = "text-ada-001"
	TextBabbage001Engine = "text-babbage-001"
	TextCurie001Engine   = "text-curie-001"
	TextDavinci001Engine = "text-davinci-001"
	TextDavinci002Engine = "text-davinci-002"
	AdaEngine            = "ada"
	BabbageEngine        = "babbage"
	CurieEngine          = "curie"
	DavinciEngine        = "davinci"
)

const (
	DefaultEngine   = Gpt35EngineTurbo
	DefaultRoleName = "User"
	// DefaultClientName TODO: load UserName to environment variable
	DefaultClientName = "mentat-analysis"
)

type EmbeddingEngine string

// TODO: add more engines released on the late 2022
const (
	TextSimilarityAda001      = "text-similarity-ada-001"
	TextSimilarityBabbage001  = "text-similarity-babbage-001"
	TextSimilarityCurie001    = "text-similarity-curie-001"
	TextSimilarityDavinci001  = "text-similarity-davinci-001"
	TextSearchAdaDoc001       = "text-search-ada-doc-001"
	TextSearchAdaQuery001     = "text-search-ada-query-001"
	TextSearchBabbageDoc001   = "text-search-babbage-doc-001"
	TextSearchBabbageQuery001 = "text-search-babbage-query-001"
	TextSearchCurieDoc001     = "text-search-curie-doc-001"
	TextSearchCurieQuery001   = "text-search-curie-query-001"
	TextSearchDavinciDoc001   = "text-search-davinci-doc-001"
	TextSearchDavinciQuery001 = "text-search-davinci-query-001"
	CodeSearchAdaCode001      = "code-search-ada-code-001"
	CodeSearchAdaText001      = "code-search-ada-text-001"
	CodeSearchBabbageCode001  = "code-search-babbage-code-001"
	CodeSearchBabbageText001  = "code-search-babbage-text-001"
	TextEmbeddingAda002       = "text-embedding-ada-002"
)

const (
	DefaultBaseURL        = "https://api.openai.com/v1"
	DefaultUserAgent      = "mentat" // go-gpt3
	DefaultTimeoutSeconds = 300
)

func getEngineURL(engine string) string {
	return fmt.Sprintf("%s/engines/%s/completions", DefaultBaseURL, engine)
}
