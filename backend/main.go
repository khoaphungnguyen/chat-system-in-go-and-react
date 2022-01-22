package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define an upgrader wih a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Need to check the original of the connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Define a reader which listens for new message to the websocket endpoint
func reader(conn *websocket.Conn) {
	for {
		// read message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print the message
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// Define the websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// Upgrade this connection to websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen infinite messages go through the websocket
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// map our '/ws' endpoint to the 'serveWs' function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App V0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
