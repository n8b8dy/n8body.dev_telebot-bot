package dtos

type GptMessageDTO struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptChoiceDTO struct {
	Index        int `json:"index"`
	Message      GptMessageDTO
	FinishReason string `json:"finish_reason"`
}

type GptUsageDTO struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GptRequestDTO struct {
	Model    string          `json:"model"`
	Messages []GptMessageDTO `json:"messages"`
	User     string          `json:"user"`
}

type GptResponseDTO struct {
	Id      string         `json:"id"`
	Object  string         `json:"object"`
	Created int            `json:"created"`
	Model   string         `json:"model"`
	Choices []GptChoiceDTO `json:"choices"`
	Usage   GptUsageDTO    `json:"usage"`
}
