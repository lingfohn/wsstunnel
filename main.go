package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	wssConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		msgType, msg, err := wssConn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("%s sent: %s \n", wssConn.RemoteAddr(), string(msg))
		if err = wssConn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":32123", nil)
}
