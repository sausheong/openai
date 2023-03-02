package openai

const (
	COMPLETIONS_URL = "https://api.openai.com/v1/completions"
	IMAGE_URL       = "https://api.openai.com/v1/images"
	EDITS_URL       = "https://api.openai.com/v1/edits"
	MODELS_URL      = "https://api.openai.com/v1/models"
	EMBEDDINGS_URL  = "https://api.openai.com/v1/embeddings"
	FILES_URL       = "https://api.openai.com/v1/files"
	FINETUNES_URL   = "https://api.openai.com/v1/finetunes"
	MODERATIONS_URL = "https://api.openai.com/v1/moderations"

	TEXT_GPT_35_TURBO = "gpt-3.5-turbo"
	TEXT_DAVINCI_003  = "text-davinci-003"
	TEXT_DAVINCI_002  = "text-davinci-002"
	TEXT_CURIE_001    = "text-curie-001"
	TEXT_BABBAGE_001  = "text-babbage-001"
	TEXT_ADA_001      = "text-ada-001"

	CODE_DAVINCI_002 = "code-davinci-002"
	CODE_CUSHMAN_001 = "code-cushman-001"
)

type Client struct {
	Authorization string
	Organization  string
	Request       map[string]any
}
