package socketmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // TODO: Make it configurable via the .env file
	WriteBufferSize: 1024, // TODO: Make it configurable via the .env file
}

// TODO: Use generics so that we can take string messages, that'd be nice!
type SocketMessage struct {
	EventType string                 `json:"eventType"`
	PadName   string                 `json:"padName"`
	Message   map[string]interface{} `json:"message"`
}

// Bind the websockets to the gin router
func BindSocket(router *gin.RouterGroup) {

	router.GET("/get", func(ctx *gin.Context) {
		webSocketUpgrade(ctx.Writer, ctx.Request)
	})

}

func webSocketUpgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %v\n", err)
		return
	}

	// Start listening to this socket
	for {
		// Try Read the JSON input from the socket
		_, msg, err := conn.ReadMessage()

		// Check if a close request was sent
		if errors.Is(err, websocket.ErrCloseSent) {
			break
		}

		if err != nil {
			// There has been an error reading the message
			fmt.Println("Failed to read from the socket")
			// Skip this cycle
			continue
		}

		// Init the variable
		var p SocketMessage
		// Try and parse the json
		err = json.Unmarshal([]byte(msg), &p)
		if err != nil {
			// There has been an error reading the message
			fmt.Println("Failed to parse the JSON", err)
			// Skip this cycle
			continue
		}

		// Pass the message to the proper handlers

		handleSocketMessage(p)
	}
}

// Handle the socket's message
func handleSocketMessage(msg SocketMessage) {

	// Check the type of message
	fmt.Println(msg.EventType)

}

func BroadcastMessage(padName string, message string) {

}
