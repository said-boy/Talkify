package entitiy

import "github.com/google/generative-ai-go/genai"

type Conversations struct {
	History []*Conversation `json:"history"`
}

func (c *Conversations) ToGenaiContent() []*genai.Content {
	return []*genai.Content{}
}

type Conversation struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Part struct {
	Text genai.Text `json:"text"`
}