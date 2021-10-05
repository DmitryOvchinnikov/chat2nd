package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize: socketBufferSize,
	WriteBufferSize: messageBufferSize,
}

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte

	// join is a channel for clients wishing to join the room.
	join chan *client

	// leave is a channel for clients wishing to leave the room.
	leave chan *client

	// clients holds all current clients in this room.
	clients map[*client]bool
}

// newRoom makes a new room.
func newRoom() *room {
	return &room{
		forward: make(chan[]byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}

}

func (r *room) run()  {
	for  {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	cl := &client{
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}

	r.join <- cl
	defer func() { r.leave <- cl }()
	go cl.write()
	cl.read()
}
