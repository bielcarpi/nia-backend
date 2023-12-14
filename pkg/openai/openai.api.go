package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"io"
	"nia-backend/config"
)

type Client struct {
	client *openai.Client
}

func OpenAIClientProvider() *Client {
	client := openai.NewClient(config.GetConfig().OpenAIAPIKey)
	return &Client{client: client}
}

func (ai *Client) SpeechToText(ctx context.Context, audio io.Reader) (string, error) {
	req := openai.AudioRequest{
		Model:  openai.Whisper1,
		Reader: audio,
	}
	resp, err := ai.client.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}

func (ai *Client) SendToGPT3(ctx context.Context, prompt string) (string, error) {
	req := openai.CompletionRequest{
		Model:     openai.GPT3Dot5Turbo1106,
		MaxTokens: 500,
		Prompt:    prompt,
	}

	resp, err := ai.client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Text, nil
}

func (ai *Client) TextToSpeech(text string) (string, error) {
	return "", nil
}
