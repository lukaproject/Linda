package client

import (
	"Linda/protocol/hbconn"
	"Linda/protocol/models"

	"github.com/gorilla/websocket"
)

type IClient interface {
	HeartBeat(*models.HeartBeatFromAgent) (*models.HeartBeatFromServer, error)
	HeartBeatStart(*models.HeartBeatStart) (*models.HeartBeatStartResponse, error)
}

// 请注意，这并不是一个协程安全的client
type Client struct {
	conn *websocket.Conn
}

func (c *Client) HeartBeat(agentHB *models.HeartBeatFromAgent) (serverHB *models.HeartBeatFromServer, err error) {
	if err = hbconn.WriteMessage(c.conn, agentHB); err != nil {
		return
	}
	serverHB = &models.HeartBeatFromServer{}
	if err = hbconn.ReadMessage(c.conn, serverHB); err != nil {
		return
	}
	return
}

func (c *Client) HeartBeatStart(req *models.HeartBeatStart) (resp *models.HeartBeatStartResponse, err error) {
	if err = hbconn.WriteMessage(c.conn, req); err != nil {
		return
	}
	resp = &models.HeartBeatStartResponse{}
	if err = hbconn.ReadMessage(c.conn, resp); err != nil {
		return
	}
	return
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
