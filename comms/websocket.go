package comms

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cyclo9/gotils"
	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	conn   *websocket.Conn
	url    string
	header http.Header
}

func NewWebsocketClient(url string, header http.Header) *WebsocketClient {
	return &WebsocketClient{
		url:    url,
		header: header,
	}
}

// Opens up a websocket connection
func (ws *WebsocketClient) Connect() {
	conn, _, err := websocket.DefaultDialer.Dial(ws.url, ws.header)
	gotils.HandleErr(err, "connecting to websocket")
	ws.conn = conn
}

// Write a message to the server; assumes the message is a byte array
func (ws *WebsocketClient) WriteMessage(message []byte) {
	if ws.conn == gotils.ZeroValue[*websocket.Conn]() {
		fmt.Println("Fail to write message. Not connected to server.")
		os.Exit(2)
	}

	err := ws.conn.WriteMessage(websocket.TextMessage, message)
	gotils.HandleErr(err, "sending message")
}

// Returns any response from the server as a byte array
func (ws *WebsocketClient) ReadMessage() []byte {
	if ws.conn == gotils.ZeroValue[*websocket.Conn]() {
		fmt.Println("Fail to read message. Not connected to server.")
		os.Exit(2)
	}

	_, resBytes, err := ws.conn.ReadMessage()
	gotils.HandleErr(err, "reading message")
	return resBytes
}
