package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"nia-backend/pkg/openai"
	"nia-backend/pkg/util"
)

func ProcessAudioHandler(c *gin.Context) {
	client := openai.ClientProvider()
	// TODO handle 429: exceeded current plan limits

	// Convert the stream of audio that's being received to text
	text, err := client.SpeechToText(c, c.Request.Body)
	if err != nil {
		util.SendError(c, err)
		return
	}

	// Send the text to GPT-3
	generatedText, err := client.SendToGPT3(c, text)
	if err != nil {
		util.SendError(c, err)
		return
	}

	// Convert the text to speech
	speech, err := client.TextToSpeech(c, generatedText)
	if err != nil {
		util.SendError(c, err)
		return
	}

	// Stream the speech back to the client
	err = sendSpeech(c, speech)
	if err != nil {
		util.SendError(c, err)
		return
	}
}

func sendSpeech(c *gin.Context, speech io.ReadCloser) error {
	defer speech.Close()

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Transfer-Encoding", "chunked") // Streaming the audio data to the client

	buffer := make([]byte, 1024)
	for {
		n, err := speech.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		c.Writer.Write(buffer[:n])
		c.Writer.(http.Flusher).Flush() // Important to flush the buffer
	}

	return nil
}
