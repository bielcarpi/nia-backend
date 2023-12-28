package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"nia-backend/pkg/openai"
	"os"
	"path"
)

func ProcessAudioHandler(c *gin.Context) {
	client := openai.ClientProvider()
	// TODO handle 429: exceeded current plan limits

	uploadDir := "./uploads"

	// Parse the multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // Max 10 MB file size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large or invalid form"})
		return
	}

	// Retrieve the file from the form data
	file, _, err := c.Request.FormFile("file") // "file" is the name attribute in the form
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	defer file.Close()

	// Generate a unique filename, for example using a UUID
	newFileName := uuid.New().String() + ".mp3"

	// Create the file in the designated directory
	filePath := path.Join(uploadDir, newFileName)
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create the file"})
		return
	}
	defer out.Close()

	// Copy the file content to the new file
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	text, err := client.SpeechToText(c, filePath)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	generatedText, err := client.SendToGPT3(c, text)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	speech, err := client.TextToSpeech(c, generatedText)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = sendSpeech(c, speech)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
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
