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

func ClientProvider() *Client {
	client := openai.NewClient(config.GetConfig().OpenAIAPIKey)
	return &Client{client: client}
}

func (ai *Client) SpeechToText(ctx context.Context, speech io.Reader) (string, error) {
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "speech.ogg", // This is a workaround to the fact that the API requires a file path
		Reader:   speech,
	}
	resp, err := ai.client.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}

func (ai *Client) SendToGPT3(ctx context.Context, prompt string) (string, error) {
	resp, err := ai.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (ai *Client) TextToSpeech(ctx context.Context, text string) (io.ReadCloser, error) {
	speech, err := ai.client.CreateSpeech(ctx, openai.CreateSpeechRequest{
		Model:          openai.TTSModel1,
		Input:          text,
		Voice:          openai.VoiceAlloy,
		ResponseFormat: openai.SpeechResponseFormatOpus, // We use OGG for better voice compression
	})
	if err != nil {
		return nil, err
	}

	return speech, nil
}
