package magic

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type MagicObject struct {
	ApiKey       string
	MaxTokens    int
	EngineModel  string
	SystemPrompt string
	Temperature  float32
}

func NewMagicObject(apiKey string) MagicObject {
	return MagicObject{
		ApiKey:       apiKey,
		MaxTokens:    500,
		EngineModel:  openai.GPT3Dot5Turbo,
		SystemPrompt: "You are a magic function executor that generate result according to the user's instruction. Please just give the result directly without any explanation",
		Temperature:  0.5,
	}
}

func (m *MagicObject) SetDefaultEngineModel(engineModel string) {
	if engineModel == "" {
		m.EngineModel = openai.GPT3Dot5Turbo
	} else {
		m.EngineModel = engineModel
	}
}

func (m *MagicObject) SetDefaultSystemPrompt(systemPrompt string) {
	if systemPrompt == "" {
		m.SystemPrompt = ""
	} else {
		m.SystemPrompt = systemPrompt
	}
}

func (m *MagicObject) SetDefaultTemperature(temperature float32) {
	if temperature <= 0 {
		m.Temperature = 0.1
	} else if temperature > 1 {
		m.Temperature = 1
	} else {
		m.Temperature = temperature
	}

}

func (m *MagicObject) DoMagic(ctx context.Context, prompt string, args map[string]interface{}) (string, error) {
	client := openai.NewClient(m.ApiKey)
	request := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: m.MaxTokens,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: m.SystemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: m.Temperature,
	}
	response, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil

}
