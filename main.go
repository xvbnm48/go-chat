package main

import "github.com/gorilla/websocket"

type M map[string]interface{}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string
}
type SocketResponse struct {
	From    string
	Type    string
	Message string
}
type WebSocketConnection struct {
	*websocket.Conn
	Username string
}
