package hbconn

import (
	"Linda/protocol/models"
	"errors"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// 将内容写入读取 websocket conn
// 传入时请保证 conn, v 非nil
func WriteMessage[T any](conn IWSConn, v *T) error {
	return conn.WriteMessage(websocket.BinaryMessage, models.Serialize(v))
}

// 读取 websocket conn中的内容
// 传入时请保证 conn, v 非nil
func ReadMessage[T any](conn IWSConn, v *T) error {
	msgType, body, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	models.Deserialize(body, v)
	if msgType != websocket.BinaryMessage {
		logrus.Debugf("msgType is invalid, %d", msgType)
		return errors.New("msgType is not binary")
	}
	return nil
}
