package client

import (
	"Linda/baselibs/abstractions/xctx"
	"Linda/protocol/hbconn"
	"Linda/protocol/models"

	"github.com/gorilla/websocket"
	"github.com/lukaproject/xerr"
)

type IClient interface {
	HeartBeat(*models.HeartBeatFromAgent) (*models.HeartBeatFromServer, error)
	HeartBeatStart(*models.HeartBeatStart) (*models.HeartBeatStartResponse, error)
}

// 请注意，这并不是一个协程安全的client
type Client struct {
	conn hbconn.IWSConn
}

func (c *Client) HeartBeat(agentHB *models.HeartBeatFromAgent) (serverHB *models.HeartBeatFromServer, err error) {
	err = xctx.NewErrHandleRun(func() {
		xerr.Must0(hbconn.WriteMessage(c.conn, agentHB))
		serverHB = &models.HeartBeatFromServer{}
		xerr.Must0(hbconn.ReadMessage(c.conn, serverHB))
	}).Err
	return
}

func (c *Client) HeartBeatStart(req *models.HeartBeatStart) (resp *models.HeartBeatStartResponse, err error) {
	err = xctx.NewErrHandleRun(func() {
		xerr.Must0(hbconn.WriteMessage(c.conn, req))
		resp = &models.HeartBeatStartResponse{}
		xerr.Must0(hbconn.ReadMessage(c.conn, resp))
	}).Err
	return
}

func (c *Client) Close() {
	c.conn.Close()
}

func New(url string) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return NewClientWithWSConn(conn), nil
}

func NewClientWithWSConn(conn hbconn.IWSConn) *Client {
	cli := &Client{}
	cli.conn = conn
	return cli
}
