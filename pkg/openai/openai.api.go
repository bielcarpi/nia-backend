package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"nia-backend/config"
)

type Client struct {
	client *openai.Client
}

func ClientProvider() *Client {
	client := openai.NewClient(config.GetConfig().OpenAIAPIKey)
	return &Client{client: client}
}

func (ai *Client) SpeechToText(ctx context.Context, fileName string) (string, error) {
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: fileName,
	}
	resp, err := ai.client.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}

func (ai *Client) SendToGPT3(ctx context.Context, prompt string) (string, error) {
	req := openai.CompletionRequest{
		Model:     openai.GPT4,
		MaxTokens: 500,
		Prompt:    prompt,
	}

	print(prompt)

	resp, err := ai.client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Text, nil
}

func (ai *Client) TextToSpeech(text string) (string, error) {
	return "", nil
}
