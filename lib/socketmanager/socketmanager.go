package socketmanager

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/JustKato/FreePad/lib/objects"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // TODO: Make it configurable via the .env file
	WriteBufferSize: 1024, // TODO: Make it configurable via the .env file
}

// The pad socket map caches all of the existing sockets
var padSocketMap map[string]map[string]*websocket.Conn = make(map[string]map[string]*websocket.Conn)

// TODO: Use generics so that we can take string messages, that'd be nice!
type SocketMessage struct {
	EventType string                 `json:"eventType"`
	PadName   string                 `json:"padName"`
	Message   map[string]interface{} `json:"message"`
}

// Bind the websockets to the gin router
func BindSocket(router *gin.RouterGroup) {

	router.GET("/get/:pad", func(ctx *gin.Context) {
		// Get the name of the pad to assign to this socket
		padName := ctx.Param("pad")
		// Upgrade the socket connection
		webSocketUpgrade(ctx, padName)
	})

}

func webSocketUpgrade(ctx *gin.Context, padName string) {

	for name, values := range ctx.Request.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	conn, err := wsUpgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Request.Header)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %v\n", err)
		return
	}

	// Check if we have any sockets in this padName
	if _, ok := padSocketMap[padName]; !ok {
		// Initialize a new map of sockets
		padSocketMap[padName] = make(map[string]*websocket.Conn)
	}

	// Give this socket a token
	socketToken := uuid.NewString()

	// Set the current connection at the socket Token position
	padSocketMap[padName][socketToken] = conn

	// Somone just connected
	UpdatePadStatus(padName)

	// Start listening to this socket
	for {
		// Try Read the JSON input from the socket
		_, msg, err := conn.ReadMessage()

		// Check if anything but a read limit was created
		if err != nil && !errors.Is(err, websocket.ErrReadLimit) {
			// Remove self from the cache
			delete(padSocketMap[padName], socketToken)
			// Somone just disconnected
			UpdatePadStatus(padName)
			break
		}

		if err != nil {
			// There has been an error reading the message
			fmt.Println("Failed to read from the socket but probably still connected")
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
		handleSocketMessage(p, socketToken, padName)
	}
}

// Handle the socket's message
func handleSocketMessage(msg SocketMessage, socketToken string, padName string) {

	// Check if this is a pad Update
	if msg.EventType == `padUpdate` {
		handlePadUpdate(msg, socketToken, padName)

		// Serialize the message
		serialized, err := json.Marshal(msg)
		// Check if there was an error
		if err != nil {
			fmt.Println(`Failed to broadcast the padUpdate`, err)
			// Stop the execution
			return
		}

		// Alert all the other pads other than this one.
		for k, pad := range padSocketMap[padName] {
			// Check if this is the same socket.
			if k == socketToken {
				// Skip self
				continue
			}

			// Send the message to the others.
			pad.WriteMessage(websocket.TextMessage, serialized)
		}
	}

}

func handlePadUpdate(msg SocketMessage, socketToken string, padName string) {

	// Check if the msg content is valid
	if _, ok := msg.Message[`content`]; !ok {
		fmt.Printf("Failed to update pad %s, invalid message\n", padName)
		return
	}

	// Check that the content is string
	newPadContent, ok := msg.Message[`content`].(string)
	if !ok {
		fmt.Printf("Type assertion failed for %s, invalid message\n", padName)
		return
	}

	// Get the pad
	pad := objects.GetPost(padName, false)
	// Update the pad contents
	pad.Content = newPadContent

	// Save to file
	objects.WritePost(pad)
}

// Update the current users of the pad about the amount of live viewers.
func UpdatePadStatus(padName string) {

	// Grab info about the map's key
	sockets, ok := padSocketMap[padName]
	// Check if the pad is set and has sockets connected.
	if !ok || len(sockets) < 1 {
		// Quit
		return
	}

	// Generate the message
	msg := SocketMessage{
		EventType: `statusUpdate`,
		PadName:   padName,
		Message: gin.H{
			// Send the current amount of live viewers
			"currentViewers": len(sockets),
		},
	}

	BroadcastMessage(padName, msg)

}

func BroadcastMessage(padName string, msg SocketMessage) {

	// Grab info about the map's key
	sockets, ok := padSocketMap[padName]
	// Check if the pad is set and has sockets connected.
	if !ok || len(sockets) < 1 {
		// Quit
		return
	}

	// Get all the participants of the pad group
	for _, s := range sockets {
		// Send the message to the socket
		s.WriteJSON(msg)
	}

}
