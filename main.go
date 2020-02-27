package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		for {
			messageType, msg, err := socket.ReadMessage()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Received: " + string(msg))
				if socket.WriteMessage(messageType, []byte("I'm sendint a response yaaayyy buddy")) != nil {
					fmt.Println("Error Occurred")
				}
			}
		}
	}
}
