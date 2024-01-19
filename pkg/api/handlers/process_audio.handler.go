package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"nia-backend/pkg/openai"
	"nia-backend/pkg/util"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins, adjust as necessary for production
	},
}

func ProcessAudioHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		util.SendError(c, err)
		return
	}
	defer conn.Close()

	client := openai.ClientProvider()

	// Whisper needs a stream of audio data
	reader, writer := io.Pipe()
	go func() {
		err := receiveAudioFromWebSocket(conn, writer)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}()

	// While the client is sending audio data, we are streaming it to the Whisper API
	text, err := client.SpeechToText(c, reader)
	if err != nil {
		util.SendError(c, err)
		return
	}

	// Send the text through the websocket
	err = conn.WriteMessage(websocket.TextMessage, []byte(text))
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

	// Send the text through the websocket
	err = conn.WriteMessage(websocket.TextMessage, []byte(generatedText))
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
	err = writeAudioToWebSocket(conn, speech)
	if err != nil {
		util.SendError(c, err)
		return
	}
}

func receiveAudioFromWebSocket(conn *websocket.Conn, writer *io.PipeWriter) error {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		if messageType == websocket.BinaryMessage {
			if _, err := writer.Write(message); err != nil {
				// Handle error
				break
			}
		} else if messageType == websocket.TextMessage {
			if string(message) == "END_OF_AUDIO" {
				err := writer.Close()
				if err != nil {
					fmt.Printf("error closing writer: %v", err)
				} // Close the writer when transmission ends
				break
			}
		}

		_, err = writer.Write(message)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeAudioToWebSocket(conn *websocket.Conn, speech io.ReadCloser) error {
	for {
		buf := make([]byte, 1024)
		n, err := speech.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		err = conn.WriteMessage(websocket.BinaryMessage, buf[:n])
		if err != nil {
			return err
		}
	}

	// Send a message to the client to indicate that the audio stream has ended
	err := conn.WriteMessage(websocket.TextMessage, []byte("END_OF_AUDIO"))
	if err != nil {
		return err
	}

	return nil
}
