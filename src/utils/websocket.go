/**
Copyright (C) 2022 Bluespada, MeGram

This file is under MIT software license, see the accompanying
file COPYING or http://www.opensource.org/licenses/mit-license
*/

package utils

import (
	"github.com/gofiber/websocket/v2"
)

// Websocket Client Structure
type WsClient struct {
	Uid     int `json:"uid"`
	Connect *websocket.Conn
}

// websocket room structure
type WsRoom struct {
	RoomId  int `json:"room_id"`
	Clients []WsClient
}

// Websocket Poll
var WsPolls = []WsClient{}
