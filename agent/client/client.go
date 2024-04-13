package client

import (
	"Linda/protocol/models"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// 请注意，这并不是一个协程安全的client
type Client struct {
	conn *websocket.Conn
}

func (c *Client) HeartBeat(agentHB *models.HeartBeatFromAgent) (serverHB *models.HeartBeatFromServer, err error) {
	if err = c.conn.WriteMessage(
		websocket.BinaryMessage,
		models.Serialize(agentHB),
	); err != nil {
		return
	}
	body, err := c.fetchResponse()
	if err != nil {
		return
	}
	serverHB = &models.HeartBeatFromServer{}
	models.Deserialize(body, serverHB)
	return
}

func (c *Client) HeartBeatStart(req *models.HeartBeatStart) (resp *models.HeartBeatStartResponse, err error) {
	if err = c.conn.WriteMessage(
		websocket.BinaryMessage,
		models.Serialize(req),
	); err != nil {
		return
	}
	body, err := c.fetchResponse()
	if err != nil {
		return
	}

	resp = &models.HeartBeatStartResponse{}
	models.Deserialize(body, resp)
	return
}

func (c *Client) fetchResponse() ([]byte, error) {
	msgType, body, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	if msgType != websocket.BinaryMessage {
		return nil, errors.New("msgType is not binary message")
	}
	return body, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func New(url string) (*Client, error) {
	cli := &Client{}
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		return nil, err
	}
	cli.conn = conn
	return cli, nil
}

func HealthCheck(url string) bool {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		logrus.Error(err)
		return false
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return false
	}

	return string(b) == "OK"
}
