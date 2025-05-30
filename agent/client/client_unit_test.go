package client_test

import (
	"Linda/agent/client"
	"Linda/protocol/models"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/suite"
)

type agentCentralClientUnitTestSuite struct {
	suite.Suite
}

type ReadMessageResult struct {
	messageType int
	p           []byte
	err         error
}

func (rmr *ReadMessageResult) Result() (int, []byte, error) {
	return rmr.messageType, rmr.p, rmr.err
}

type WriteMessageResult struct {
	err error
}

type testWSConn struct {
	closed             bool
	readMessageResult  *ReadMessageResult
	writeMessageResult *WriteMessageResult
}

func (ws *testWSConn) Close() error {
	if ws.closed {
		return nil
	}
	ws.closed = true
	return nil
}

func (ws *testWSConn) ReadMessage() (messageType int, p []byte, err error) {
	if ws.readMessageResult != nil {
		return ws.readMessageResult.Result()
	}
	return
}

func (ws *testWSConn) WriteMessage(messageType int, data []byte) (err error) {
	if ws.writeMessageResult != nil {
		return ws.writeMessageResult.err
	}
	return
}

func (ws *testWSConn) injectReadMessageResult(messageType int, p []byte, err error) {
	ws.readMessageResult = &ReadMessageResult{
		messageType: messageType,
		p:           p,
		err:         err,
	}
}

func (ws *testWSConn) injectWriteMessageResult(err error) {
	ws.writeMessageResult = &WriteMessageResult{
		err: err,
	}
}

func (s *agentCentralClientUnitTestSuite) TestHBStart_Success() {
	ws := &testWSConn{}
	client := client.NewClientWithWSConn(ws)
	ws.injectWriteMessageResult(nil)
	hbStartResp := &models.HeartBeatStartResponse{
		Result: models.OK,
	}
	ws.injectReadMessageResult(websocket.BinaryMessage, models.Serialize(hbStartResp), nil)
	resp, err := client.HeartBeatStart(&models.HeartBeatStart{})
	s.Nil(err)
	s.Equal(hbStartResp, resp)
}

func TestAgentCentralClientUnitTestMain(t *testing.T) {
	suite.Run(t, new(agentCentralClientUnitTestSuite))
}
