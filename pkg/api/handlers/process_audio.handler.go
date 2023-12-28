package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"nia-backend/pkg/openai"
	"nia-backend/pkg/util"
	"os"
	"path/filepath"
)

const UPLOAD_DIR = "./uploads"

func ProcessAudioHandler(c *gin.Context) {
	client := openai.ClientProvider()
	// TODO handle 429: exceeded current plan limits

	// Receive the audio from the client and temporarily store it
	filePath, err := receiveSpeech(c)
	defer os.Remove(filePath)

	// Convert the audio to text
	text, err := client.SpeechToText(c, filePath)
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

func receiveSpeech(c *gin.Context) (string, error) {
	fileName := uuid.New().String() + ".ogg"
	filePath := filepath.Join(UPLOAD_DIR, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the data from the request body to the file
	_, err = io.Copy(file, c.Request.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
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
