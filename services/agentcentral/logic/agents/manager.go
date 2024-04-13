package agents

import "github.com/gorilla/websocket"

type Mgr interface {
	NewNode(nodeId string, conn *websocket.Conn) error
}
