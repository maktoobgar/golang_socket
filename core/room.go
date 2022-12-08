// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package core

import "sync"

type Message struct {
	Message []byte

	From *Client
}

// Room maintains the set of active clients and broadcasts messages to the
// clients.
type Room struct {
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

func NewRoom() *Room {
	return &Room{
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
