// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package chat

import (
	"fmt"
	"sync"
	"time"
)

// Cleans rooms every 1 hour
func init() {
	ticker := time.NewTicker(time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				for key := range Rooms {
					Rooms[key].terminate <- true
				}
				Rooms = map[string]*Room{}
			}
		}
	}()
}

type Message struct {
	Message []byte

	From *Client
}

var Rooms = map[string]*Room{}

// Room maintains the set of active clients and broadcasts messages to the
// clients.
type Room struct {
	// Unique key for every room
	ID int

	// Assigned name for room
	Name string

	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Terminate room
	terminate chan bool

	// Key to lock when editing clients and broadcasting message
	clientsKey *sync.Mutex
}

func NewRoom(name string, id int) *Room {
	return &Room{
		ID:         id,
		Name:       name,
		Clients:    make(map[*Client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		terminate:  make(chan bool),
		clientsKey: &sync.Mutex{},
	}
}

func (h *Room) Run() {
	for {
		select {
		case client := <-h.register:
			h.wrapByLockUnlock(func() { h.Clients[client] = true })
		case client := <-h.unregister:
			h.wrapByLockUnlock(func() { delete(h.Clients, client) })
		case message := <-h.broadcast:
			h.wrapByLockUnlock(func() {
				message.Message = []byte(fmt.Sprintf("%s:%s", message.From.name, message.Message))
				for client := range h.Clients {
					select {
					case client.send <- message:
					default:
						delete(h.Clients, client)
					}
				}
			})
		case <-h.terminate:
			h.wrapByLockUnlock(func() {
				for client := range h.Clients {
					client.Terminate()
				}
				h.Clients = map[*Client]bool{}
			})
			return
		}
	}
}

func (h *Room) wrapByLockUnlock(function func()) {
	h.clientsKey.Lock()
	function()
	h.clientsKey.Unlock()
}

func (h *Room) Terminate() {
	h.terminate <- true
}
